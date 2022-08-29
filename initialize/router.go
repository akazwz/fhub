package initialize

import (
	"github.com/akazwz/fhub/api"
	"github.com/akazwz/fhub/api/auth"
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

	return r
}
