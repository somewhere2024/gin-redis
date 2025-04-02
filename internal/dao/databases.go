package dao

import (
	"gin-redis/internal/models"
	"gin-redis/internal/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	databaseURL := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(databaseURL), &gorm.Config{}) // &gorm.Config{}参数为可选， 用来给gorm设置一些配置
	if err != nil {
		service.Logger.Error("The database connection failed")

	}
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		service.Logger.Error("The database migration failed")
		panic("failed to migrate database")
	}
	return
}
