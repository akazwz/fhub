package service

import (
	"github.com/akazwz/fhub/global"
	"github.com/akazwz/fhub/model"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type FileService struct{}

func (s *FileService) CreateFile(file model.File, fileUri model.FileURI) error {
	file.ID = gonanoid.Must(32)
	fileUri.ID = gonanoid.Must(32)
	// transaction
	err := global.GDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&file).Error
		if err != nil {
			return err
		}
		err = tx.Create(&fileUri).Error
		return err
	})
	return err
}

func (s *FileService) CreateFilePre(file model.File) error {
	file.ID = gonanoid.Must(32)
	err := global.GDB.Create(&file).Error
	return err
}

func (s *FileService) FindFileURIBySha256(sha256 string) (string, error) {
	var fileURI model.FileURI
	err := global.GDB.Where("sha256 = ?", sha256).First(&fileURI).Error
	if err != nil {
		return "", err
	}
	return fileURI.Provider, nil
}

func (s *FileService) CreateFolder(file model.File) error {
	file.ID = gonanoid.Must(32)
	err := global.GDB.Create(&file).Error
	return err
}

func (s *FileService) FindFiles(uid, prefixDir string) ([]model.File, error) {
	var fileList []model.File
	err := global.GDB.Where("uid = ? and prefix_dir = ?", uid, prefixDir).Find(&fileList).Error
	return fileList, err
}

func (s *FileService) FindFileURI(id, uid string) (string, error) {
	var file model.File
	err := global.GDB.Where("id = ? and uid = ?", id, uid).First(&file).Error
	if err != nil {
		return "", err
	}
	var fileUri model.FileURI
	err = global.GDB.Where("sha256 = ?", file.SHA256).First(&fileUri).Error
	if err != nil {
		return "", err
	}

	return fileUri.Provider, nil
}

func (s *FileService) FindFileByUidPrefixName(uid, prefixDir, name string) {
	global.GDB.Where("uid = ? and prefix_dir = ? and name = ?", uid, prefixDir, name)
}
