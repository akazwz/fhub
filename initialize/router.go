package initialize

import (
	"github.com/akazwz/fhub/api"
	"github.com/akazwz/fhub/api/auth"
	"github.com/akazwz/fhub/api/file"
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

	fileGroup := r.Group("/file").Use(middleware.JWTAuth())
	{
		// 创建文件
		fileGroup.POST("/", file.CreateFile)
		// 创建文件夹
		fileGroup.POST("/folder", file.CreateFolder)
		// 查询文件列表
		fileGroup.GET("/file/list", file.FindFilesByParentID)
		// 查询文件路径
		fileGroup.GET("/file/path", file.FindFilesByParentID)
		//
		fileGroup.GET("/folder/:id/folders", file.FindFoldersByParentID)
		fileGroup.GET("/folder/:id/all", file.FindFoldersAndFilesByParentID)
		fileGroup.GET("/keywords/:keywords", file.FindFilesByKeywords)
		fileGroup.GET("/:id/uri", file.FindFileURI)
	}

	return r
}
