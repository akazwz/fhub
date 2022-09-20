package file

import (
	"github.com/akazwz/fhub/model"
	"github.com/akazwz/fhub/model/request"
	"github.com/akazwz/fhub/model/response"
	"github.com/akazwz/fhub/service"
	"github.com/gin-gonic/gin"
)

var fileService = service.FileService{}

// CreateFile 新建文件
func CreateFile(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	var f request.File
	err := c.ShouldBind(&f)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}

	// TODO category
	category := ""
	fileExtension := ""

	file := model.File{
		Category:      category,
		Name:          f.Name,
		ParentID:      f.ParentID,
		Size:          int64(f.Size),
		ContentHash:   f.ContentHash,
		FileExtension: fileExtension,
		UID:           uid,
	}

	provider := model.Provider{
		ContentHash: f.ContentHash,
		Provider:    f.Provider,
		URI:         f.URI,
	}

	err = fileService.CreateFile(file, provider)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
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

// FindFileURI 获取文件 uri
func FindFileURI(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	id := c.Param("id")
	file, err := fileService.FindFileByUIDAndID(uid, id)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	uri, err := fileService.FindURIByHash(file.ContentHash)
	if err != nil {
		return
	}
	response.Ok(200, uri, "success", c)
}
