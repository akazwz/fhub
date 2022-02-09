package model

import uuid "github.com/satori/go.uuid"

type User struct {
	Model
	UID      uuid.UUID `json:"uid" gorm:"unique comment: 用户 uid"`
	Username string    `json:"username" gorm:"unique comment: 用户名"`
	Password string    `json:"password" gorm:"unique comment: 用户密码"`
	Email    string    `json:"email" gorm:"unique comment: 邮箱"`
	Phone    string    `json:"phone" gorm:"unique comment: 手机号"`
	Role     int       `json:"role" gorm:"default:user comment: 角色"`
	Gender   int       `json:"gender" gorm:"comment: 性别"`
	Avatar   string    `json:"avatar" gorm:"comment: 头像"`
}

func (u User) TableName() string {
	return "user"
}
