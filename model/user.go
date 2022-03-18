package model

import (
	"github.com/akazwz/fhub/utils/crypt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key;"`
	UID       uuid.UUID `json:"uid" gorm:"not null;unique;type:varchar(255);comment: 用户uid;"`
	Username  string    `json:"username" gorm:"not null;unique;type:varchar(255);comment:用户名;"`
	Password  string    `json:"password" gorm:"not null;type:varchar(64);comment:加密后密码;"`
	Email     string    `json:"email" gorm:"unique;type:varchar(255);default:null;comment:Email;"`
	Phone     string    `json:"phone" gorm:"unique;type:varchar(255);default:null;comment:手机号;"`
	Role      int       `json:"role" gorm:"type:tinyint;default:1;comment:角色;"`
	Gender    int       `json:"gender" gorm:"type:tinyint;default:null;comment:性别;"`
	Avatar    string    `json:"avatar" gorm:"comment:头像;"`
	CreatedAt int       `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt int       `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}

func (u User) TableName() string {
	return "user"
}

// BeforeCreate hooks 新增用户前 加密密码
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UID = uuid.NewV4()
	err, hashedPwd := crypt.HashAndSortPwd(u.Password)
	if err != nil {
		return
	}
	u.Password = hashedPwd

	return
}
