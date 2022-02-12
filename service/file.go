package service

import (
	"github.com/akazwz/gin/global"
	"github.com/akazwz/gin/model"
	"github.com/akazwz/gin/model/request"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func UploadFileService(uploadFile request.UploadFile, uid uuid.UUID) (err error) {
	fileUri := model.FileURI{
		SHA256: uploadFile.Sha256,
		QKey:   uploadFile.QKey,
		CID:    uploadFile.CID,
	}

	file := model.File{
		File:      uploadFile.File,
		FileName:  uploadFile.Filename,
		PrefixDir: uploadFile.PrefixDir,
		Size:      uploadFile.Size,
		SHA256:    uploadFile.Sha256,
		UID:       uid,
	}

	/* transaction */
	err = global.GDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&fileUri).Error
		if err != nil {
			return err
		}
		err = tx.Create(&file).Error
		if err != nil {
			return err
		}
		return nil
	})
	return
}
