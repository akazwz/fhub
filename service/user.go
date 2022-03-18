package service

import (
	"errors"

	"github.com/akazwz/fhub/global"
	"github.com/akazwz/fhub/model"
	"github.com/akazwz/fhub/model/request"
	"github.com/akazwz/fhub/utils/crypt"
)

// RegisterByUsernamePwdService 用户名密码注册
func RegisterByUsernamePwdService(u request.RegisterByUsernamePwd) (err error) {
	err = global.GDB.Create(&model.User{
		Username: u.Username,
		Password: u.Password,
	}).Error
	return
}

// LoginByUsernamePwdService 用户名密码登录
func LoginByUsernamePwdService(u request.LoginByUsernamePwd) (err error, userInst *model.User) {
	var user model.User

	/* 查找用户出错 */
	err = global.GDB.Where("username = ?", u.Username).
		First(&user).Error
	if err != nil {
		return err, nil
	}
	/* 检查密码 */
	isCheck := crypt.ComparePassword(user.Password, u.Password)
	if !isCheck {
		err = errors.New("password is not right")
	}
	return err, &user
}

func GetUserProfileByUID(uid string) (err error, userInst model.User) {
	err = global.GDB.Where("uid = ?", uid).First(&userInst).Error
	return
}
