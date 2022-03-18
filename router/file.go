package router

import (
	v1 "github.com/akazwz/fhub/api/v1"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/file/uptoken", v1.GetUploadFileToken)
	routerGroup.POST("/file", v1.CreateFile)
	routerGroup.POST("/file/folder", v1.CreateFolder)
	routerGroup.GET("/file", v1.GetFileList)
	routerGroup.GET("/file/uri", v1.GetFileURI)
}
