package config

import (
	"gorm.io/driver/mysql"
    "gorm.io/gorm"
)
var DB *gorm.DB

func ConnectDB() *gorm.DB{
	dsn := "root@tcp(127.0.0.1:3306)/go-test?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	return db
}