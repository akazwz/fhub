package model

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Folder struct {
	Model
	Name     string `gorm:"unique_index:idx_only_one_name"`
	UID      string `json:"uid"`
	ParentID string `json:"parent_id"`
	Starred  bool   `json:"starred"`

	Paths []string `gorm:"-"`
}

func (folder *Folder) TableName() string {
	return "folders"
}

func (folder *Folder) BeforeCreate(*gorm.DB) (err error) {
	folder.ID = gonanoid.MustGenerate(alphabet, 32)
	return
}

// Create 创建文件夹
func (folder *Folder) Create(db *gorm.DB) error {
	return db.Create(folder).Error
}

// GetChildFolder 获取子文件夹
func (folder *Folder) GetChildFolder(db *gorm.DB) ([]Folder, error) {
	folders := make([]Folder, 0)
	result := db.Where("parent_file_id = ?", folder.ID).Find(&folders)
	return folders, result.Error
}

// GetPaths 获取路径
func (folder *Folder) GetPaths(db *gorm.DB) error {
	if folder.ParentID == "root" {
		folder.Paths = append(folder.Paths, "root")
		return nil
	}
	var parentFolder Folder
	err := db.
		Where("id = ? AND uid = ?", folder.ParentID, folder.UID).
		First(&parentFolder).Error

	if err == nil {
		err := parentFolder.GetPaths(db)
		folder.Paths = append(folder.Paths, parentFolder.Name)
		return err
	}
	return err
}
