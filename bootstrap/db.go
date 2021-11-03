package bootstrap

import (
	"gorm.io/gorm"
	"jasy/goblog/app/models/article"
	"jasy/goblog/app/models/user"
	"jasy/goblog/pkg/model"
	"time"
)

func SetupDB() {
	db := model.ConnectDB()
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	migration(db)
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&user.User{}, &article.Article{})
}