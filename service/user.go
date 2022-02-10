package service

import (
	"github.com/akazwz/gin/global"
	"github.com/akazwz/gin/model"
	"github.com/akazwz/gin/model/request"
)

func RegisterByUsernamePwdService(u request.RegisterByUsernamePwd) (err error) {
	err = global.GDB.Create(&model.User{
		Username: u.Username,
		Password: u.Password,
	}).Error
	return
}
