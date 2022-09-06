package service

import (
	"github.com/akazwz/fhub/global"
	"github.com/akazwz/fhub/model"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

var alphabet = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

type ObjectService struct{}

func (s *ObjectService) CreateFile(file model.File, fileUri model.FileURI) error {
	file.ID = gonanoid.MustGenerate(alphabet, 32)
	fileUri.ID = gonanoid.Must(32)
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

func (s *ObjectService) FindFileByUIDAndID(uid, id string) (model.File, error) {
	file, err := model.FindFileByUIDAndID(global.GDB, uid, id)
	return file, err
}

func (s *ObjectService) FindFileURIBySha256(sha256 string) (string, error) {
	var fileURI model.FileURI
	provider, err := fileURI.FindFileURIBySha256(global.GDB, sha256)
	if err != nil {
		return "", err
	}
	return provider, nil
}

func (s *ObjectService) CreateFolder(folder model.Folder) error {
	err := folder.Create(global.GDB)
	return err
}

func (s *ObjectService) FindFilesByParentID(uid, parentID string) ([]model.File, error) {
	files, err := model.GetFilesByParentID(global.GDB, uid, parentID)
	if err != nil {
		return nil, err
	}
	return files, err
}

func (s *ObjectService) FindFilesByKeywords(uid string, parents []string, keywords string) ([]model.File, error) {
	files, err := model.GetFilesByKeywords(global.GDB, uid, parents, keywords)
	if err != nil {
		return nil, err
	}
	return files, err
}
