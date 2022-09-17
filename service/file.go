package service

import (
	"github.com/akazwz/fhub/global"
	"github.com/akazwz/fhub/model"
	"gorm.io/gorm"
)

type FileService struct{}

func (s *FileService) CreateFile(file model.File, fileUri model.FileURI) error {
	// transaction
	err := global.GDB.Transaction(func(tx *gorm.DB) error {
		// create file
		err := file.Create(tx)
		if err != nil {
			return err
		}
		err = fileUri.Create(tx)
		return err
	})
	return err
}

func (s *FileService) FindFileByUIDAndID(uid, id string) (model.File, error) {
	file, err := model.FindFileByUIDAndID(global.GDB, uid, id)
	return file, err
}

func (s *FileService) FindFileURIBySha256(sha256 string) (string, error) {
	var fileURI model.FileURI
	provider, err := fileURI.FindFileURIBySha256(global.GDB, sha256)
	if err != nil {
		return "", err
	}
	return provider, nil
}

func (s *FileService) CreateFolder(folder model.Folder) error {
	err := folder.Create(global.GDB)
	return err
}

func (s *FileService) FindFilesByParentID(uid, parentID string) ([]model.File, error) {
	files, err := model.FindFilesByParentID(global.GDB, uid, parentID)
	if err != nil {
		return nil, err
	}
	return files, err
}

func (s *FileService) FindFilesByKeywords(uid string, parents []string, keywords string) ([]model.File, error) {
	files, err := model.FindFilesByKeywords(global.GDB, uid, parents, keywords)
	if err != nil {
		return nil, err
	}
	return files, err
}
