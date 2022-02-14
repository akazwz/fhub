package v1

import (
	"os"

	"github.com/akazwz/gin/global"
	"github.com/akazwz/gin/model"
	"github.com/akazwz/gin/model/request"
	"github.com/akazwz/gin/model/response"
	"github.com/akazwz/gin/service"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// GetUploadFileToken 获取文件上传token(qiniu)
func GetUploadFileToken(c *gin.Context) {
	accessKey := global.CONF.Qiniu.AccessKey
	if len(os.Getenv("QAK")) > 0 {
		accessKey = os.Getenv("QAK")
	}

	secretKey := global.CONF.Qiniu.SecretKey
	if len(os.Getenv("QSK")) > 0 {
		secretKey = os.Getenv("QSK")
	}

	mac := qbox.NewMac(accessKey, secretKey)
	bucket := "akazwz"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)

	type TokenUpload struct {
		Token string `json:"token"`
	}

	response.Ok(CodeSuccessCreateUploadToken, TokenUpload{
		Token: upToken,
	}, "获取成功", c)
}

// CreateFile 保存文件信息到数据库
func CreateFile(c *gin.Context) {
	/* get uid */
	claims, _ := c.Get("claims")
	customClaims := claims.(*model.MyCustomClaims)
	userUID := customClaims.UID

	var file request.UploadFile

	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.BadRequest(CodeErrorBindJson, "参数错误", c)
		return
	}

	if len(file.CID) < 1 && len(file.QKey) < 1 {
		response.BadRequest(CodeErrorBindJson, "qkey 和 cid 不能都为空", c)
		return
	}

	err = service.UploadFileService(file, userUID)
	if err != nil {
		response.BadRequest(CodeErrorCreatFile, "上传失败", c)
		return
	}
	response.Created(CodeSuccessCreateFile, nil, "上传成功", c)
}

// GetFileList 获取文件列表
func GetFileList(c *gin.Context) {
	/* get uid */
	claims, _ := c.Get("claims")
	customClaims := claims.(*model.MyCustomClaims)
	userUID := customClaims.UID

	/* 获取文件路径前缀 */
	prefixDir := c.Query("prefix_dir")

	if len(prefixDir) < 1 {
		response.BadRequest(CodeErrorParams, "参数错误", c)
		return
	}

	err, fileList := service.GetFileListService(userUID, prefixDir)
	if err != nil {
		response.BadRequest(CodeErrorGetFileList, "获取文件列表失败", c)
		return
	}

	response.Ok(CodeSuccessGetFileList, fileList, "获取成功", c)
}
