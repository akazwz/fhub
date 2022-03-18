package v1

import (
	"github.com/akazwz/fhub/model/response"
	"github.com/gin-gonic/gin"
)

// DownloadByMagnet 根据magnet 下载
func DownloadByMagnet(c *gin.Context) {
	/* 获取文件Key */
	magnet := c.Query("magnet")

	response.Ok(CodeSuccessGetFileUri, gin.H{"uri": magnet}, "获取成功", c)
}
