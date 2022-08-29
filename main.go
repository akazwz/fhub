package main

import (
	"log"
	"net/http"
	"os"

	"github.com/akazwz/fhub/initialize"
	"github.com/joho/godotenv"
)

func main() {
	//counter.Consume()
	// 读取环境变量配置
	InitEnvConfig()
	// 初始化 gorm db
	initialize.InitGormDB()

	// 迁移表
	initialize.MigrateTables()

	// 初始化路由
	r := initialize.InitRouter()
	// 端口地址
	port := os.Getenv("API_PORT")
	s := &http.Server{
		Addr:    port,
		Handler: r,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
		log.Fatalln("Api启动失败")
	}
}

// InitEnvConfig 读取 env 配置文件
func InitEnvConfig() {
	// 非生产环境读取配置文件
	if os.Getenv("MODE") != "prod" {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatalln("读取配置文件失败")
		}
	}
}
