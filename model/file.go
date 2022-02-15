package model

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type File struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	FID           uuid.UUID `json:"fid" gorm:"column:fid;not null;unique;type:varchar(255);comment: 文件fid;"`
	File          bool      `json:"file" gorm:"not null;comment: 是否是文件;"`
	FileName      string    `json:"file_name" gorm:"not null;comment: 文件名;"`
	PrefixDir     string    `json:"prefix_dir" gorm:"not null;default:0/;comment:文件前缀路径;"`
	UIDUniqueFile string    `json:"uid_unique_file" gorm:"not null;unique;comment:uid-文件路径前缀文件名;"`
	Size          int64     `json:"size" gorm:"not null;default:0;comment:文件大小;"`
	SHA256        string    `json:"sha256" gorm:"column:sha256;not null;type:varchar(255);comment:文件sha256;"`
	UID           uuid.UUID `json:"uid" gorm:"not null;type:varchar(255);comment:用户uid;"`
	CreatedAt     int       `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt     int       `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}

func (f File) TableName() string {
	return "file"
}

// BeforeCreate hooks
func (f *File) BeforeCreate(tx *gorm.DB) (err error) {
	f.FID = uuid.NewV4()
	f.UIDUniqueFile = fmt.Sprintf("%s-%s-%s", f.UID, f.PrefixDir, f.FileName)
	return
}

type FileURI struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	SHA256    string `json:"sha256" gorm:"column:sha256;not null;unique;type:varchar(255);comment:文件sha256;"`
	QKey      string `json:"qkey" gorm:"column:qkey;type:varchar(255);comment:七牛key;"`
	CID       string `json:"cid" gorm:"column:cid;type:varchar(255);comment:ipfs cid;"`
	CreatedAt int    `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt int    `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}

func (f FileURI) TableName() string {
	return "file_uri"
}

// BeforeCreate hooks
func (f *FileURI) BeforeCreate(tx *gorm.DB) (err error) {
	return
}
