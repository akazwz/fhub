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

func GetUploadFileToken(c *gin.Context) {
	accessKey := global.CONF.Qiniu.AccessKey
	if len(os.Getenv("QAK")) > 0 {
		accessKey = os.Getenv("QAK")
	}

	secretKey := global.CONF.Qiniu.SecretKey
	if len(os.Getenv("QSK")) > 0 {
		accessKey = os.Getenv("QSK")
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
