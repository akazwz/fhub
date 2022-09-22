package main

import (
	"github.com/getsentry/sentry-go"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/akazwz/fhub/initialize"
	"github.com/joho/godotenv"
)

func init() {
	// 读取环境变量配置
	InitEnvConfig()
	// 初始化 gorm db
	initialize.InitGormDB()
	// 迁移表
	initialize.MigrateTables()

	initialize.InitR2Client()

	initialize.InitWasabiClient()

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)
}

func main() {
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
