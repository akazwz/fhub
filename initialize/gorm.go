package initialize

import (
	"fmt"
	"github.com/akazwz/gin/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

// InitDB  初始化数据库
func InitDB() *gorm.DB {
	/* 配置文件中配置 */
	dbConfig := global.CONF.DataBase

	/* 环境变量覆盖配置文件 */
	dbUser := dbConfig.User
	if len(os.Getenv("DB_USER")) > 0 {
		dbUser = os.Getenv("DB_USER")
	}
	dbPassword := dbConfig.Password
	if len(os.Getenv("DB_PASSWORD")) > 0 {
		dbPassword = os.Getenv("DB_PASSWORD")
	}
	dbHost := dbConfig.Host
	if len(os.Getenv("DB_HOST")) > 0 {
		dbHost = os.Getenv("DB_HOST")
	}

	dbName := dbConfig.Name
	if len(os.Getenv("DB_NAME")) > 0 {
		dbName = os.Getenv("DB_NAME")
	}

	/* 获取 dsn */
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		/* 数据库连接失败 */
		return nil
	}
	return db
}

// CreateTables 数据库表迁移
func CreateTables(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		log.Fatal("数据库建表失败")
	}
}
