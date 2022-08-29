package auth

import (
	"github.com/akazwz/fhub/model"
	"github.com/akazwz/fhub/model/request"
	"github.com/akazwz/fhub/model/response"
	"github.com/akazwz/fhub/service"
	"github.com/gin-gonic/gin"
)

// SignupByUsernamePwd 注册(用户名 + 密码)
func SignupByUsernamePwd(c *gin.Context) {
	authService := service.AuthService{}
	// 获取绑定参数
	signup := request.SignUp{}
	err := c.ShouldBind(&signup)
	if err != nil {
		response.BadRequest(400, nil, "参数错误", c)
		return
	}
	// 构造用户
	user := model.User{
		Username: signup.Username,
		Password: signup.Password,
	}
	// 注册
	userInstance, err := authService.SignupService(user)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Created(200, userInstance, "注册成功", c)
}
