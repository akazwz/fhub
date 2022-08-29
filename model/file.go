package model

type File struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Type      string `json:"type" gorm:"not null;"`
	Name      string `json:"name" gorm:"not null;"`
	PrefixDir string `json:"prefix_dir" gorm:"not null;"`
	Size      int    `json:"size" gorm:"not null;default:0;"`
	SHA256    string `json:"sha256" gorm:"column:sha256;not null;type:varchar(255);"`
	UID       string `json:"uid" gorm:"not null;type:varchar(255);"`
	CreatedAt int    `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt int    `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}

func (f File) TableName() string {
	return "files"
}

type FileURI struct {
	ID        string `json:"id" gorm:"primary_key"`
	SHA256    string `json:"sha256" gorm:"column:sha256;not null;unique;type:varchar(255);"`
	Providers string `json:"providers" gorm:"not null;"`
	CreatedAt int    `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt int    `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}

func (f FileURI) TableName() string {
	return "file_uris"
}
