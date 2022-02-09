package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello, %s!", name),
		})
	})

	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Println("run error")
		return
	}
}
