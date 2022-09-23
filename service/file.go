package service

import (
	"github.com/akazwz/fhub/global"
	"github.com/akazwz/fhub/model"
	"gorm.io/gorm"
)

type FileService struct{}

func (s *FileService) CreateFileAndProvider(file model.File, provider model.Provider) error {
	// transaction
	err := global.GDB.Transaction(func(tx *gorm.DB) error {
		// 创建文件
		err := file.Create(tx)
		if err != nil {
			return err
		}
		// 创建provider
		err = provider.Create(tx)
		// TODO 更新用户空间
		return err
	})
	return err
}

func (s *FileService) FindFileByUIDAndID(uid, id string) (model.File, error) {
	file, err := model.FindFileByUIDAndID(global.GDB, uid, id)
	return file, err
}

func (s *FileService) FindFileByUIDParentIdAndName(uid, parentId, name string) (model.File, error) {
	file, err := model.FindFileByUIDParentIdAndName(global.GDB, uid, parentId, name)
	return file, err
}

func (s *FileService) FindProviderByContentHash(hash string) (*model.Provider, error) {
	provider := model.Provider{ContentHash: hash}
	err := provider.FindProviderByContentHash(global.GDB)
	if err != nil {
		return nil, err
	}
	return &provider, nil
}

func (s *FileService) FindFilesByKeywords(uid string, parents []string, keywords string) ([]model.File, error) {
	files, err := model.FindFilesByKeywords(global.GDB, uid, parents, keywords)
	if err != nil {
		return nil, err
	}
	return files, err
}

func (s *FileService) DeleteFileByUIDAndID(uid, id string) error {
	return model.DeleteFileByUIDAndId(global.GDB, uid, id)
}
