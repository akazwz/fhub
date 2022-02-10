package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/akazwz/gin/global"
	"github.com/akazwz/gin/initialize"
)

func main() {
	global.VP = initialize.InitViper()

	if global.VP == nil {
		log.Println("初始化配置失败")
	}

	global.GDB = initialize.InitDB()

	if global.GDB != nil {
		initialize.CreateTables(global.GDB)
		db, _ := global.GDB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
	} else {
		log.Println("数据库连接失败")
		return
	}

	routers := initialize.Routers()

	/* 端口地址 */
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	s := &http.Server{
		Addr:    ":" + port,
		Handler: routers,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Println("系统启动失败")
	}
}
