package model

import uuid "github.com/satori/go.uuid"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key;"`
	UID       uuid.UUID `json:"uid" gorm:"not null;unique;type:varchar(32);comment: 用户uid;"`
	Username  string    `json:"username" gorm:"not null;unique;type:varchar(255);comment:用户名;"`
	Password  string    `json:"password" gorm:"not null;type:varchar(64);comment:加密后密码;"`
	Email     string    `json:"email" gorm:"unique;type:varchar(255);comment:Email;"`
	Phone     string    `json:"phone" gorm:"unique;type:varchar(255);comment:手机号;"`
	Role      int       `json:"role" gorm:"type:tinyint;default:1;comment:角色;"`
	Gender    int       `json:"gender" gorm:"type:tinyint;comment:性别;"`
	Avatar    string    `json:"avatar" gorm:"comment:头像;"`
	CreatedAt int       `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt int       `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}

func (u User) TableName() string {
	return "user"
}
