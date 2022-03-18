package router

import (
	v1 "github.com/akazwz/fhub/api/v1/user"
	"github.com/gin-gonic/gin"
)

func InitPublicRouter(routerGroup *gin.RouterGroup) {
	/* 注册用户 */
	routerGroup.POST("/user", v1.CreateUserByUsernamePwd)
	/* 登录获取token */
	routerGroup.POST("/user/token", v1.CreateTokenByUsernamePwd)
}
