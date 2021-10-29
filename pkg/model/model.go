package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jasy/goblog/pkg/logger"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error

	config := mysql.New(mysql.Config{
		DSN: "root:w*$l#1tv4OSIOoJ1lxi%Eau4Ob5WuDIW@tcp(127.0.0.1:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})

	DB, err = gorm.Open(config, &gorm.Config{})

	logger.LogError(err)

	return DB
}