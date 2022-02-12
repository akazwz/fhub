package router

import (
	v1 "github.com/akazwz/gin/api/v1"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/file", v1.CreateFile)
}
