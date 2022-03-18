package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() *gorm.DB {
	var err error
	DB, err = gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")))
	if err != nil {
		panic(err)
	}
	return DB
}
