package model

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type FileURI struct {
	ID        string `json:"id" gorm:"primary_key"`
	SHA256    string `json:"sha256" gorm:"column:sha256;not null;type:varchar(255);"`
	Provider  string `json:"provider" gorm:"not null;"`
	CreatedAt int    `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt int    `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}

func (fileURI *FileURI) TableName() string {
	return "file_uris"
}

func (fileURI *FileURI) BeforeCreate(*gorm.DB) (err error) {
	fileURI.ID = gonanoid.MustGenerate(alphabet, 32)
	return
}

func (fileURI *FileURI) Create(db *gorm.DB) error {
	return db.Create(fileURI).Error
}

func (fileURI *FileURI) FindFileURIBySha256(db *gorm.DB, sha256 string) (string, error) {
	err := db.Where("sha256 = ?", sha256).First(&fileURI).Error
	if err != nil {
		return "", err
	}
	return fileURI.Provider, nil
}
