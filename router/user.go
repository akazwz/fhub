package router

import (
	v1 "github.com/akazwz/gin/api/v1/user"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/user/profile", v1.GetUserProfileByToken)
}
