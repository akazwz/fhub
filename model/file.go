package model

import (
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

var alphabet = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

type File struct {
	Model
	Category      string `json:"category"`
	ContentHash   string `json:"content_hash"`
	UID           string `json:"uid"`
	ParentID      string `json:"parent_id"`
	FileExtension string `json:"file_extension"`
	MimeType      string `json:"mime_type"`
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	Starred       bool   `json:"starred"`
	Thumbnail     string `json:"thumbnail"`
}

func (file *File) TableName() string {
	return "files"
}

func (file *File) BeforeCreate(*gorm.DB) (err error) {
	file.ID = gonanoid.MustGenerate(alphabet, 32)
	return
}

// Create 创建文件
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

// FindFileByUIDAndID 通过 uid 和 id 查找文件
func FindFileByUIDAndID(db *gorm.DB, uid, id string) (File, error) {
	var file File
	err := db.Where("uid = ? AND id = ?", uid, id).Find(&file).Error
	return file, err
}

// Rename 重命名文件
func (file *File) Rename(db *gorm.DB, newName string) error {
	return db.Model(&file).UpdateColumn("name", newName).Error
}

// FindFilesByKeywords 通过关键词搜索文件
func FindFilesByKeywords(db *gorm.DB, uid string, parents []string, keywords ...interface{}) ([]File, error) {
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

// FindFilesByParentID 通过 parent id查找文件
func FindFilesByParentID(db *gorm.DB, uid, id string) ([]File, error) {
	files := make([]File, 0)
	result := db.Where("uid = ? and parent_id = ?", uid, id).Find(&files)
	return files, result.Error
}
