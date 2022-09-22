package file

import (
	"github.com/akazwz/fhub/model"
	"github.com/akazwz/fhub/model/request"
	"github.com/akazwz/fhub/model/response"
	"github.com/akazwz/fhub/service"
	"github.com/akazwz/fhub/utils"
	"github.com/akazwz/fhub/utils/s3/wasabi"
	"github.com/gin-gonic/gin"
)

var fileService = service.FileService{}

// CreateFile 新建文件
func CreateFile(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	var f request.CreateFile
	err := c.ShouldBind(&f)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}

	file := model.File{
		ContentHash: f.ContentHash,
		UID:         uid,
		ParentID:    f.ParentID,
		Name:        f.Name,
		Size:        f.Size,
	}
	// 查找 provider
	provider, err := fileService.FindProviderByContentHash(f.ContentHash)
	// hash 匹配直接创建文件
	if err == nil {
		err := fileService.CreateFile(file, *provider)
		if err != nil {
			response.BadRequest(400, nil, err.Error(), c)
			return
		}
		response.Ok(200, file, "success", c)
		return
	}

	// hash 不匹配, 返回 multipart upload url
	key := utils.GenerateID(32)
	upload := wasabi.CreateMultipartUpload(key)
	uploadId := upload.UploadId
	uploadPartUrlList := make(map[int32]interface{}, 0)

	for _, part := range f.PartInfoList {
		uploadPart := wasabi.CreatePresignUploadPart(*uploadId, key, part.PartNumber)
		uploadPartUrlList[part.PartNumber] = uploadPart.URL
	}

	response.Ok(200, gin.H{
		"upload_id":       upload,
		"upload_url_list": uploadPartUrlList,
	}, "success", c)
}

func GetPathByParentID() {

}

func FindFoldersByParentID(c *gin.Context) {

}

func FindFoldersAndFilesByParentID(c *gin.Context) {

}

func FindFilesByKeywords(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	keywords := c.Param("keywords")
	files, err := fileService.FindFilesByKeywords(uid, nil, keywords)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, files, "success", c)
}

func FindFoldersByKeywords() {

}
