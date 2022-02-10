package router

import (
	v1 "github.com/akazwz/gin/api/v1/user"
	"github.com/gin-gonic/gin"
)

func InitPublicRouter(routerGroup *gin.RouterGroup) {
	/* 注册用户 */
	routerGroup.POST("/user", v1.CreateUserByUsernamePwd)
}
