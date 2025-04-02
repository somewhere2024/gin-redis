package main

import (
	"gin-redis/internal/api"
	"gin-redis/internal/dao"
	"gin-redis/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.InitDB()
	service.LogInit()
	app := gin.Default()
	app.GET("/status", api.StatusTest)
	app.POST("/login", api.LoginUser)
	app.POST("/register", api.RegisterUser)
	app.GET("/readme", api.ReadMe)
	app.Run(":8000")
}
