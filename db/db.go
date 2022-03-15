package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() *gorm.DB {
	var err error
	DB, err = gorm.Open(mysql.Open("root:123456789@tcp(127.0.0.1:3306)/healthy?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	return DB
}
