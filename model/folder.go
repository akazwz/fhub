package model

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type FolderPath struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// Folder name uid parentID 联合唯一索引
type Folder struct {
	Model
	Name     string `json:"name" gorm:"size:191;uniqueIndex:idx_name"`
	UID      string `json:"-" gorm:"size:191;uniqueIndex:idx_name"`
	ParentID string `json:"parent_id" gorm:"size:191;uniqueIndex:idx_name"`
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

// BeforeDelete 文件夹删除时，文件夹内的文件夹以及文件一同删除
func (folder *Folder) BeforeDelete(db *gorm.DB) error {
	// 查找文件夹下文件
	files, err := folder.GetChildFiles(db)
	if err != nil {
		return err
	}

	if len(files) > 0 {
		// 获取文件 id 列表
		fileIDList := make([]string, 0)
		for _, file := range files {
			fileIDList = append(fileIDList, file.ID)
		}
		// 根据文件 id 列表删除文件
		err = DeleteFilesByIDList(db, fileIDList)
		if err != nil {
			return err
		}
	}

	// 查找子文件夹
	folders, err := folder.GetChildFolders(db)
	if err != nil {
		return err
	}

	if len(folders) > 0 {
		for _, folder := range folders {
			// delete 才能出发 BeforeDelete
			err = folder.Delete(db)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Create 创建文件夹
func (folder *Folder) Create(db *gorm.DB) error {
	return db.Create(folder).Error
}

// Delete 删除文件夹
func (folder *Folder) Delete(db *gorm.DB) error {
	err := db.
		Where("uid = ? AND id = ?", folder.UID, folder.ID).
		Delete(folder).
		Error
	if err != nil {
		return err
	}
	return nil
}

// Rename 重命名文件夹
func (folder *Folder) Rename(db *gorm.DB) error {
	err := db.
		Model(folder).
		Where("uid = ? AND id = ?", folder.UID, folder.ID).
		Update("name", folder.Name).
		Error
	if err != nil {
		return err
	}
	return nil
}

// GetChildFolders 获取子文件夹
func (folder *Folder) GetChildFolders(db *gorm.DB) ([]Folder, error) {
	folders := make([]Folder, 0)
	result := db.Where("uid = ? AND parent_id = ?", folder.UID, folder.ID).
		Order("created_at").
		Find(&folders)
	return folders, result.Error
}

// GetChildFiles 获取子文件
func (folder *Folder) GetChildFiles(db *gorm.DB) ([]File, error) {
	files := make([]File, 0)
	result := db.Where("uid = ? AND parent_id = ?", folder.UID, folder.ID).
		Order("created_at").
		Find(&files)
	return files, result.Error
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

func DeleteFoldersByIDList(db *gorm.DB, idList []string) error {
	folder := Folder{}
	return db.Where("id IN ?", idList).Delete(&folder).Error
}
