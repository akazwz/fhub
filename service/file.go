package service

import (
	"github.com/akazwz/gin/global"
	"github.com/akazwz/gin/model"
	"github.com/akazwz/gin/model/request"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// UploadFileService 保存文件信息到数据库
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

// CreateFolderService 保存文件夹信息到数据库
func CreateFolderService(folder request.NewFolder, uid uuid.UUID) (err error) {
	file := model.File{
		File:      false,
		FileName:  folder.FolderName,
		PrefixDir: folder.PrefixDir,
		Size:      0,
		UID:       uid,
	}

	err = global.GDB.Create(&file).Error
	return
}

// GetFileListService 根据文件前缀和uid获取文件列表
func GetFileListService(uid uuid.UUID, prefixDir string) (err error, fileList []model.File) {
	err = global.GDB.Where("uid = ? AND prefix_dir = ?", uid, prefixDir).Find(&fileList).Error
	return
}

// GetFileQKeyByFID 根据 uid 和 fid 获取文件QKey
func GetFileQKeyByFID(uid uuid.UUID, fid string) (err error, QKey string) {
	file := model.File{}
	err = global.GDB.Where("uid = ? AND fid = ?", uid, fid).Find(&file).Error
	if err != nil {
		return
	}
	fileUri := model.FileURI{}
	err = global.GDB.Where("sha256 = ?", file.SHA256).Find(&fileUri).Error
	if err != nil {
		return
	}
	return nil, fileUri.QKey
}
