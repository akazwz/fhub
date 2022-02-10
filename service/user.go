package service

import (
	"errors"

	"github.com/akazwz/gin/global"
	"github.com/akazwz/gin/model"
	"github.com/akazwz/gin/model/request"
	"github.com/akazwz/gin/utils/crypt"
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
