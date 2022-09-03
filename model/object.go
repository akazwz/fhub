package model

type Object struct {
	ID        string `json:"id" gorm:"primary_key"`
	Type      string `json:"type" gorm:"not null;"`
	Name      string `json:"name" gorm:"not null;index:idx_only_one_name"`
	ParentID  string `json:"parent_id" gorm:"not null;index:idx_only_one_name"`
	Size      int    `json:"size" gorm:"not null;default:0;"`
	SHA256    string `json:"sha256" gorm:"column:sha256;not null;type:varchar(255);"`
	UID       string `json:"uid" gorm:"not null;type:varchar(255);"`
	CreatedAt int    `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt int    `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}

func (o *Object) TableName() string {
	return "objects"
}

type ObjectURI struct {
	ID        string `json:"id" gorm:"primary_key"`
	SHA256    string `json:"sha256" gorm:"column:sha256;not null;type:varchar(255);"`
	Provider  string `json:"provider" gorm:"not null;"`
	CreatedAt int    `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt int    `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}

func (o *ObjectURI) TableName() string {
	return "objects_uris"
}
