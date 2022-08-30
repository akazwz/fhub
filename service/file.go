package service

import (
	"github.com/akazwz/fhub/global"
	"github.com/akazwz/fhub/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type FileService struct{}

func (s *FileService) CreateFile(file model.File, fileUri model.FileURI) error {
	file.ID = uuid.NewV4().String()
	fileUri.ID = uuid.NewV4().String()
	err := global.GDB.Transaction(func(tx *gorm.DB) error {
		err := global.GDB.Create(&file).Error
		if err != nil {
			return err
		}
		err = global.GDB.Create(&fileUri).Error
		return err
	})
	return err
}

func (s *FileService) CreateFolder(file model.File) error {
	file.ID = uuid.NewV4().String()
	err := global.GDB.Create(&file).Error
	return err
}

func (s *FileService) FindFiles(uid, prefixDir string) ([]model.File, error) {
	var fileList []model.File
	err := global.GDB.Where("uid = ? and prefix_dir = ?", uid, prefixDir).Find(&fileList).Error
	return fileList, err
}
