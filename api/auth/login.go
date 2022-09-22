package auth

import (
	"github.com/akazwz/fhub/model"
	"github.com/akazwz/fhub/model/request"
	"github.com/akazwz/fhub/model/response"
	"github.com/akazwz/fhub/service"
	"github.com/akazwz/fhub/utils"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

// LoginByUsernamePwd 登录(用户名 + 密码)
func LoginByUsernamePwd(c *gin.Context) {
	authService := service.AuthService{}
	// 获取绑定参数
	login := request.Login{}
	err := c.ShouldBind(&login)
	if err != nil {
		sentry.CaptureException(err)
		response.BadRequest(400, nil, "参数错误", c)
		return
	}
	// 构造用户
	user := model.User{
		Username: login.Username,
		Password: login.Password,
	}
	// 登录
	userInstance, err := authService.LoginService(user)
	if err != nil {
		sentry.CaptureException(err)
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	TokenNext(*userInstance, c)
}

func TokenNext(user model.User, c *gin.Context) {
	j := utils.NewJWT()
	claims := j.CreateClaims(request.BaseClaims{UID: user.ID})
	token, err := j.CreateToken(claims)
	if err != nil {
		response.BadRequest(400, nil, "获取token失败", c)
		return
	}
	response.Created(400, gin.H{
		"user":       user,
		"token":      token,
		"expires_at": claims.RegisteredClaims.ExpiresAt,
	}, "登录成功", c)
}
