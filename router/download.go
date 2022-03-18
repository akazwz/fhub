package router

import (
	v1 "github.com/akazwz/fhub/api/v1"
	"github.com/gin-gonic/gin"
)

func InitDownloadRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/download/magnet", v1.DownloadByMagnet)
}
