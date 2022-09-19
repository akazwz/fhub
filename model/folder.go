package model

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type FolderPath struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Folder struct {
	Model
	Name     string `gorm:"unique_index:idx_only_one_name" json:"name"`
	UID      string `json:"-"`
	ParentID string `json:"parent_id"`
	Starred  bool   `json:"starred"`

	Path []FolderPath `gorm:"-" json:"-"`
}

func (folder *Folder) TableName() string {
	return "folders"
}

func (folder *Folder) BeforeCreate(*gorm.DB) error {
	id, err := gonanoid.Generate(alphabet, 32)
	folder.ID = id
	return err
}

// Create 创建文件夹
func (folder *Folder) Create(db *gorm.DB) error {
	return db.Create(folder).Error
}

// GetChildFolders 获取子文件夹
func (folder *Folder) GetChildFolders(db *gorm.DB) ([]Folder, error) {
	folders := make([]Folder, 0)
	result := db.Where("uid = ? AND parent_id = ?", folder.UID, folder.ID).Find(&folders)
	return folders, result.Error
}

// GetPath 获取路径
func (folder *Folder) GetPath(db *gorm.DB) error {
	if folder.ParentID == "root" {
		folder.Path = append(folder.Path, FolderPath{
			Name: "Root",
			ID:   "root",
		})
		return nil
	}
	var parentFolder Folder
	err := db.
		Where("uid = ? AND id = ?", folder.UID, folder.ParentID).
		First(&parentFolder).Error

	if err == nil {
		folder.Path = append(folder.Path, FolderPath{
			Name: parentFolder.Name,
			ID:   parentFolder.ID,
		})
		// 变更 parentID
		folder.ParentID = parentFolder.ParentID
		err := folder.GetPath(db)
		return err
	}
	return err
}
