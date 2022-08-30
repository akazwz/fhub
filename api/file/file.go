package file

import (
	"github.com/akazwz/fhub/model"
	"github.com/akazwz/fhub/model/request"
	"github.com/akazwz/fhub/model/response"
	"github.com/akazwz/fhub/service"
	"github.com/gin-gonic/gin"
)

var fileService = service.FileService{}

func CreateFilePre(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	var f request.PreFile
	err := c.ShouldBind(&f)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}

	uri, err := fileService.FindFileURIBySha256(f.SHA256)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	file := model.File{
		Type:      "file",
		Name:      f.Name,
		PrefixDir: f.PrefixDir,
		Size:      f.Size,
		SHA256:    f.SHA256,
		UID:       uid,
	}

	err = fileService.CreateFilePre(file)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, uri, "success", c)
}

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
		Type:      "file",
		Name:      f.Name,
		PrefixDir: f.PrefixDir,
		Size:      f.Size,
		SHA256:    f.SHA256,
		UID:       uid,
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

func CreateFolder(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	var folder request.Folder
	err := c.ShouldBind(&folder)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}

	file := model.File{
		Type:      "folder",
		Name:      folder.Name,
		PrefixDir: folder.PrefixDir,
		Size:      0,
		SHA256:    "",
		UID:       uid,
	}

	err = fileService.CreateFolder(file)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Created(200, nil, "success", c)
}

func FindFiles(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	prefixDir := c.GetHeader("prefix-dir")

	files, err := fileService.FindFiles(uid, prefixDir)
	if err != nil {
		response.BadRequest(300, nil, err.Error(), c)
		return
	}
	response.Ok(200, files, "success", c)
}

func FindFileURI(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	fid := c.Param("id")
	fileProvider, err := fileService.FindFileURI(fid, uid)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, fileProvider, "success", c)
}
