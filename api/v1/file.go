package v1

import (
	"os"
	"time"

	"github.com/akazwz/fhub/global"
	"github.com/akazwz/fhub/model"
	"github.com/akazwz/fhub/model/request"
	"github.com/akazwz/fhub/model/response"
	"github.com/akazwz/fhub/service"
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

// CreateFile 上传文件
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

// CreateFolder 新建文件夹
func CreateFolder(c *gin.Context) {
	/* get uid */
	claims, _ := c.Get("claims")
	customClaims := claims.(*model.MyCustomClaims)
	userUID := customClaims.UID

	var folder request.NewFolder

	err := c.ShouldBindJSON(&folder)
	if err != nil {
		response.BadRequest(CodeErrorBindJson, "参数错误", c)
		return
	}

	err = service.CreateFolderService(folder, userUID)
	if err != nil {
		response.BadRequest(CodeErrorCreateFolder, "新建文件夹失败", c)
		return
	}
	response.Created(CodeSuccessCreateFolder, nil, "新建成功", c)
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

// GetFileURI 获取文件 uri
func GetFileURI(c *gin.Context) {
	/* get uid */
	claims, _ := c.Get("claims")
	customClaims := claims.(*model.MyCustomClaims)
	userUID := customClaims.UID

	/* 获取文件Key */
	fid := c.Query("fid")

	err, QKey := service.GetFileQKeyByFID(userUID, fid)
	if err != nil {
		response.BadRequest(CodeErrorGetFileUri, "获取文件URI", c)
		return
	}

	if len(QKey) < 1 {
		response.BadRequest(CodeErrorParams, "参数错误", c)
		return
	}

	accessKey := global.CONF.Qiniu.AccessKey
	if len(os.Getenv("QAK")) > 0 {
		accessKey = os.Getenv("QAK")
	}

	secretKey := global.CONF.Qiniu.SecretKey
	if len(os.Getenv("QSK")) > 0 {
		secretKey = os.Getenv("QSK")
	}

	mac := qbox.NewMac(accessKey, secretKey)
	domain := "https://file.hellozwz.com"

	deadline := time.Now().Add(time.Second * 3600).Unix() // 1h
	privateAccessURL := storage.MakePrivateURLv2(mac, domain, QKey, deadline)
	if len(privateAccessURL) < 1 {
		response.BadRequest(CodeErrorGetFileUri, "获取失败， uri为空", c)
		return
	}

	response.Ok(CodeSuccessGetFileUri, gin.H{"uri": privateAccessURL}, "获取成功", c)
}
