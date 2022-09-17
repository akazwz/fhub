package model

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID        string `json:"id" gorm:"primary_key;"`
	Username  string `json:"username" gorm:"not null;unique;type:varchar(255);"`
	Password  string `json:"-" gorm:"not null;type:varchar(64);"`
	Email     string `json:"email" gorm:"unique;type:varchar(255);default:null;"`
	Phone     string `json:"phone" gorm:"unique;type:varchar(255);default:null;"`
	Role      int    `json:"role" gorm:"type:tinyint;default:1;"`
	CreatedAt int    `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt int    `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(*gorm.DB) error {
	id, err := gonanoid.Generate(alphabet, 32)
	u.ID = id
	return err
}
