package v1

import (
	"github.com/akazwz/gin/model"
	"github.com/akazwz/gin/model/request"
	"github.com/akazwz/gin/model/response"
	"github.com/akazwz/gin/service"
	"github.com/gin-gonic/gin"
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
