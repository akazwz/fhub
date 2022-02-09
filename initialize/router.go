package initialize

import (
	"github.com/akazwz/gin/model/response"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Routers 路由
func Routers() *gin.Engine {
	router := gin.Default()

	/* cors 使用跨域中间件 */
	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))

	/* 404 路由不存在 */
	router.NoRoute(response.NotFound)

	//Teapot  418
	router.GET("teapot", func(c *gin.Context) {
		c.JSON(http.StatusTeapot, gin.H{
			"message": "I'm a teapot",
			"story": "This code was defined in 1998 " +
				"as one of the traditional IETF April Fools' jokes," +
				" in RFC 2324, Hyper Text Coffee Pot Control Protocol," +
				" and is not expected to be implemented by actual HTTP servers." +
				" However, known implementations do exist.",
		})
	})

	return router
}
