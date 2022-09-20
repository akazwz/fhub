package initialize

import (
	"github.com/akazwz/fhub/api"
	"github.com/akazwz/fhub/api/auth"
	"github.com/akazwz/fhub/api/file"
	"github.com/akazwz/fhub/api/folder"
	"github.com/akazwz/fhub/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	/* cors */
	r.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))

	r.NoRoute(api.NotFound)
	r.GET("/healthz", api.Healthz)
	r.GET("", api.Endpoints)
	r.GET("teapot", api.Teapot)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", auth.SignupByUsernamePwd)
		authGroup.POST("/login", auth.LoginByUsernamePwd)
		authGroup.GET("/me", middleware.JWTAuth(), auth.Me)
	}

	folderGroup := r.Group("/folders").Use(middleware.JWTAuth())
	{
		folderGroup.GET("/:id/path", folder.FindPath)
		folderGroup.GET("/:id/all", folder.FindFoldersAndFilesByParentID)
		folderGroup.GET("/:id/folders", folder.FindFoldersByParentID)
		folderGroup.GET("/:id/files", folder.FindFilesByParentID)
		folderGroup.POST("/:id/:name", folder.CreateFolder)
		folderGroup.DELETE("/:id", folder.DeleteFolder)
		folderGroup.PATCH("/:id/name/:name", folder.RenameFolder)
	}

	fileGroup := r.Group("/files").Use(middleware.JWTAuth())
	{
		fileGroup.GET("/", file.CreateFile)
		fileGroup.GET("/:id/uri", file.FindFileURI)
	}

	return r
}
