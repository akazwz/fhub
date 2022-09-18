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

	file := model.File{
		Name:        f.Name,
		ParentID:    f.ParentID,
		Size:        int64(f.Size),
		ContentHash: f.SHA256,
		UID:         uid,
	}

	fileUri := model.FileURI{
		SHA256:   f.SHA256,
		Provider: f.Provider,
	}

	err = fileService.CreateFile(file, fileUri)
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

func FindFileURI(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	id := c.Param("id")
	provider, err := fileService.FindFileByUIDAndID(uid, id)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, provider, "success", c)
}
