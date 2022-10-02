package file

import (
	"fmt"
	"os"
	"time"

	"github.com/akazwz/fhub/global"
	"github.com/akazwz/fhub/model"
	"github.com/akazwz/fhub/model/request"
	"github.com/akazwz/fhub/model/response"
	"github.com/akazwz/fhub/service"
	"github.com/akazwz/fhub/utils"
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

	category := utils.GetCategoryByName(f.Name)
	name := f.Name
	fileByName, err := fileService.FindFileByUIDParentIdAndName(uid, f.ParentID, f.Name)
	// 文件名存在
	if fileByName.Name == f.Name {
		ext := utils.GetExtByName(name)
		name = utils.GetPureNameByName(name)
		if len(ext) > 0 {
			name = fmt.Sprintf("%s-%s.%s", name, time.Now().Format("2006-01-02 15:04:05"), ext)
		} else {
			name = fmt.Sprintf("%s-%s", name, time.Now().Format("2006-01-02 15:04:05"))
		}
	}

	file := model.File{
		ContentHash: f.ContentHash,
		UID:         uid,
		ParentID:    f.ParentID,
		Name:        name,
		Size:        f.Size,
		Category:    category,
	}
	// 查找 provider
	provider, err := fileService.FindProviderByContentHash(f.ContentHash)
	// hash 匹配直接创建文件
	if err == nil {
		err := fileService.CreateFileAndProvider(file, *provider)
		if err != nil {
			response.BadRequest(400, nil, err.Error(), c)
			return
		}
		response.Ok(200, file, "success", c)
		return
	}

	// hash 不匹配, 返回 multipart upload url
	accessKey := os.Getenv("AK")
	accessSecret := os.Getenv("SK")
	url := os.Getenv("QINIU_URL")
	client, err := utils.S3Storage.NewS3Client(accessKey, accessSecret, url)
	bucket := "fhub"

	key := utils.GenerateID(32)

	upload := utils.S3Storage.CreateMultipartUpload(client, bucket, key)
	uploadId := upload.UploadId
	uploadPartUrlList := make(map[int32]interface{}, 0)

	for _, part := range f.PartInfoList {
		uploadPart := utils.S3Storage.CreatePresignUploadPart(client, bucket, key, *uploadId, part.PartNumber)
		uploadPartUrlList[part.PartNumber] = uploadPart.URL
	}

	response.Ok(200, gin.H{
		"upload_id":       upload,
		"upload_url_list": uploadPartUrlList,
	}, "success", c)
}

func CompleteMultipartUpload(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	var complete request.CompleteMultipartUpload
	err := c.ShouldBind(&complete)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}

	name := complete.Name
	size := complete.Size
	key := complete.Key
	uploadId := complete.UploadId
	contentHash := complete.ContentHash
	parentId := complete.ParentId

	accessKey := os.Getenv("AK")
	accessSecret := os.Getenv("SK")
	url := os.Getenv("QINIU_URL")
	client, err := utils.S3Storage.NewS3Client(accessKey, accessSecret, url)
	bucket := "fhub"

	_, err = utils.S3Storage.CompleteUpload(client, bucket, key, uploadId, contentHash)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	category := utils.GetCategoryByName(name)

	fileByName, err := fileService.FindFileByUIDParentIdAndName(uid, parentId, name)
	// 文件名存在
	if fileByName.Name == name {
		ext := utils.GetExtByName(name)
		name = utils.GetPureNameByName(name)
		if len(ext) > 0 {
			name = fmt.Sprintf("%s-%s.%s", name, time.Now().Format("2006-01-02 15:04:05"), ext)
		} else {
			name = fmt.Sprintf("%s-%s", name, time.Now().Format("2006-01-02 15:04:05"))
		}
	}
	// 文件
	file := model.File{
		ContentHash: contentHash,
		UID:         uid,
		ParentID:    parentId,
		Name:        name,
		Size:        size,
		Category:    category,
	}
	// provider  wasabi
	provider := model.Provider{
		ContentHash: contentHash,
		Provider:    "wasabi",
		Key:         key,
	}

	err = fileService.CreateFileAndProvider(file, provider)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, file, "success", c)
}

func DeleteFileByID(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	id := c.Param("id")
	err := fileService.DeleteFileByUIDAndID(uid, id)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, nil, "success", c)
}

func ListMultipartUpload(c *gin.Context) {
	key := c.Query("key")
	uploadId := c.Query("upload_id")
	client := global.R2Client
	bucket := os.Getenv("R2_BUCKET_NAME")
	output, err := utils.S3Storage.FindUploadPart(client, bucket, key, uploadId)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, output, "success", c)
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
