package service

import (
	"errors"

	"github.com/akazwz/fhub/global"
	"github.com/akazwz/fhub/model"
	"github.com/akazwz/fhub/utils"
	"gorm.io/gorm"
)

type AuthService struct{}

func (authService *AuthService) SignupService(u model.User) (*model.User, error) {
	user := authService.FindUserByUsername(u.Username)
	// username 已经存在
	if user != nil {
		return user, errors.New("用户名已注册")
	}
	u.Password = utils.BcryptHash(u.Password)

	err := global.GDB.Transaction(func(tx *gorm.DB) error {
		// 创建用户
		err := global.GDB.Create(&u).Error
		if err != nil {
			return err
		}
		// 创建容量
		capacity := model.Capacity{
			UID:   u.ID,
			Total: 10 * 1024 * 1024 * 1024,
			Used:  0,
			Album: 0,
			Video: 0,
		}
		err = capacity.Create(tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &u, err
}

func (authService *AuthService) LoginService(u model.User) (*model.User, error) {
	user := authService.FindUserByUsername(u.Username)
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	isPasswordCorrect := utils.BcryptCheck(u.Password, user.Password)
	if !isPasswordCorrect {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

func (authService *AuthService) FindUserByUsername(username string) *model.User {
	var user model.User
	err := global.GDB.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func (authService *AuthService) FindUserByID(id string) *model.User {
	var user model.User
	err := global.GDB.Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func (authService *AuthService) FindCapacityByUID(uid string) *model.Capacity {
	var cap model.Capacity
	err := global.GDB.Where("uid = ?", uid).First(&cap).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &cap
}
