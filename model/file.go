package model

import (
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

var alphabet = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

type File struct {
	Model
	Name     string `gorm:"unique_index:idx_only_one"`
	Size     uint64
	SHA256   string `json:"sha256" gorm:"column:sha256;type:varchar(255);"`
	ParentID string `gorm:"index:parent_id;unique_index:idx_only_one"`
	UID      string `gorm:"index:user_id;unique_index:idx_only_one"`

	Position string `gorm:"-"`
}

func (file *File) TableName() string {
	return "files"
}

func (file *File) BeforeCreate(*gorm.DB) (err error) {
	file.ID = gonanoid.MustGenerate(alphabet, 32)
	return
}

func (file *File) Create(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		// create file
		if err := tx.Create(file).Error; err != nil {
			log.Println("create file error")
			return err
		}
		// TODO update user storage
		return nil
	})
	return err
}

func FindFileByUIDAndID(db *gorm.DB, uid, id string) (File, error) {
	var file File
	err := db.Where("uid = ? AND id = ?", uid, id).Find(&file).Error
	return file, err
}

func (file *File) Rename(db *gorm.DB, newName string) error {
	return db.Model(&file).UpdateColumn("name", newName).Error
}

func GetFilesByKeywords(db *gorm.DB, uid string, parents []string, keywords ...interface{}) ([]File, error) {
	var (
		conditions string
		files      []File
		result     = db
	)

	// conditions
	for i := 0; i < len(keywords); i++ {
		conditions += "name like ?"
		if i != len(keywords)-1 {
			conditions += " or "
		}
	}

	// uid
	if uid != "" {
		result = result.Where("uid = ?", uid)
	}

	// parents
	if len(parents) > 0 {
		result = result.Where("parents_id IN (?)", parents)
	}

	// result
	result = result.Where("("+conditions+")", keywords...).Find(&files)

	return files, result.Error
}

func GetFilesByParentID(db *gorm.DB, uid, id string) ([]File, error) {
	files := make([]File, 0)
	result := db.Where("uid = ? and parent_id = ?", uid, id).Find(&files)
	return files, result.Error
}
