package initialize

import (
	"fmt"
	"log"
	"os"

	"github.com/akazwz/fhub/global"
	"github.com/akazwz/fhub/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGormDB() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_NAME"),
	)
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatalln("Init Gorm DB Error")
	} else {
		global.GDB = db
	}
}

func MigrateTables() {
	if err := global.GDB.AutoMigrate(
		model.User{},
		model.File{},
		model.FileURI{},
		model.Folder{},
	); err != nil {
		log.Fatalln(err)
	}
}
