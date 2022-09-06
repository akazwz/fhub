package model

import (
	"gorm.io/gorm"
	"path"
)

type Folder struct {
	Model
	Name     string `gorm:"unique_index:idx_only_one_name"`
	ParentID string `gorm:"index:parent_id;unique_index:idx_only_one_name"`
	UID      string `gorm:"index:uid"`

	Position string `gorm:"-"`
}

func (folder *Folder) Create(db *gorm.DB) error {
	return db.Create(folder).Error
}

func (folder *Folder) GetChildFolder(db *gorm.DB) ([]Folder, error) {
	folders := make([]Folder, 0)
	result := db.Where("parent_id = ?", folder.ID).Find(&folders)
	if result.Error == nil {
		for i := 0; i < len(folders); i++ {
			folders[i].Position = path.Join(folder.Position, folder.Name)
		}
	}
	return folders, result.Error
}

func (folder *Folder) TraceRoot(db *gorm.DB) error {
	if folder.ParentID == "root" {
		return nil
	}
	var parentFolder Folder
	err := db.
		Where("id = ? AND uid = ?", folder.ParentID, folder.UID).
		First(&parentFolder).Error

	if err == nil {
		err := parentFolder.TraceRoot(db)
		folder.Position = path.Join(parentFolder.Position, parentFolder.Name)
		return err
	}
	return err
}
