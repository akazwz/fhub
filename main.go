package main

import (
	"github.com/akazwz/gin/global"
	"github.com/akazwz/gin/initialize"

	"database/sql"
	"log"
	"net/http"
	"os"
)

func main() {
	/* 初始化配置 */
	global.VP = initialize.InitViper()

	if global.VP == nil {
		log.Println("初始化配置失败")
	}

	/* 初始化数据库  */
	global.DB = initialize.InitDB()

	if global.DB != nil {
		/* 建表 */
		initialize.CreateTables(global.DB)
		db, _ := global.DB.DB()
		/* 关闭数据库连接 */
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
	} else {
		log.Println("数据库连接失败")
	}

	routers := initialize.Routers()

	/* 端口地址 */
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	s := &http.Server{
		Addr:    port,
		Handler: routers,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Println("系统启动失败")
	}
}
