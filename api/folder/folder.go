package folder

import (
	"github.com/akazwz/fhub/model"
	"github.com/akazwz/fhub/model/response"
	"github.com/akazwz/fhub/service"
	"github.com/gin-gonic/gin"
)

var folderService = service.FolderService{}

// CreateFolder 新建文件夹
func CreateFolder(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	parentID := c.Param("id")
	folderName := c.Param("name")

	folder := model.Folder{
		Name:     folderName,
		ParentID: parentID,
		UID:      uid,
	}

	err := folderService.CreateFolder(folder)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Created(200, nil, "success", c)
}

// DeleteFolder 删除文件夹
func DeleteFolder(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	id := c.Param("id")

	folder := model.Folder{
		Model: model.Model{ID: id},
		UID:   uid,
	}

	err := folderService.DeleteFolder(folder)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, folder, "success", c)
}

// RenameFolder 重命名文件夹
func RenameFolder(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	id := c.Param("id")
	folderName := c.Param("name")

	folder := model.Folder{
		Model: model.Model{ID: id},
		UID:   uid,
		Name:  folderName,
	}

	err := folderService.RenameFolder(folder)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, folder, "success", c)
}

// FindFilesByParentID 根据 parent id查找文件
func FindFilesByParentID(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	parentID := c.Param("id")
	files, err := folderService.FindFilesByParentID(uid, parentID)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, files, "success", c)
}

// FindFoldersByParentID 根据 parent id查找子文件夹
func FindFoldersByParentID(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	parentID := c.Param("id")
	folders, err := folderService.FindFoldersByParentID(uid, parentID)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, folders, "success", c)
}

// FindFoldersAndFilesByParentID 根据 parent id查找文件和子文件夹
func FindFoldersAndFilesByParentID(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	parentID := c.Param("id")
	files, err := folderService.FindFilesByParentID(uid, parentID)
	folders, err := folderService.FindFoldersByParentID(uid, parentID)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, gin.H{
		"files":   files,
		"folders": folders,
	}, "success", c)
}

// FindPath 根据 parent 查找路径
func FindPath(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	parentID := c.Param("id")

	path, err := folderService.FindPath(uid, parentID)
	if err != nil {
		response.BadRequest(400, nil, err.Error(), c)
		return
	}
	response.Ok(200, path, "success", c)
}
