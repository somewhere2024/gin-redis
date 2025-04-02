package api

import (
	"gin-redis/internal/dao"
	"gin-redis/internal/models"
	"gin-redis/internal/service"
	"gin-redis/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"net/http"
)

func StatusTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "运行成功"})
}

func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名或密码不能为空"})
		return
	}
	UUid := service.CreateUUId()
	dao.DB.Create(&models.User{
		Username: username,
		Password: password,
		UUId:     UUid,
	})
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func LoginUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		service.Logger.Info("用户名或密码为空")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名或密码不能为空"})
		return
	}
	user := &models.User{}
	// 用redis做缓存 (未实现)
	result := dao.DB.Where("username = ?", username).First(user)
	if result.Error != nil {
		service.Logger.Info("用户名或密码错误")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名或密码错误"})
		return
	}
	ok := service.VerifyPassword(password, user)
	if !ok {
		service.Logger.Info("用户名或密码错误")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名或密码错误"})
		return
	}
	// 创建token返回给用户
	token, err := service.CreateToken(jwt.MapClaims{
		"username": user.Username,
		"uuid":     user.UUId,
	})
	if err != nil {
		service.Logger.Error("创建token失败", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器错误"})
		return
	}
	c.Header("YourToken", token)
	c.JSON(http.StatusOK, gin.H{"message": "登录成功"})

}

// 获取自己的信息
// 并用redis做缓存机制
func ReadMe(c *gin.Context) {
	token := c.GetHeader("Authorization")
	c.JSON(http.StatusOK, gin.H{"message": "获取成功", "access_token": token, "token_type": "Bearer"})
	return
}

func RedisSet(c *gin.Context) {
	key := c.PostForm("key")
	value := c.PostForm("value")
	if key == "" || value == "" {
		service.Logger.Error("redis测试失败:key和value不能为空")
		c.JSON(http.StatusBadRequest, gin.H{"message": "redis测试失败"})
		return
	}

	err := utils.SetKeyValue(dao.RDB, key, value)
	if err != nil {
		service.Logger.Error("redis测试失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"message": "redis测试失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "redis测试成功"})
}

func RedisGet(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		service.Logger.Error("redis测试失败:key不能为空")
		c.JSON(http.StatusBadRequest, gin.H{"message": "key不能为空"})
		return
	}
	value, err := utils.GetKeyValue(dao.RDB, key)
	if value == "" && err != nil {
		service.Logger.Error("redis测试失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"message": "redis测试失败"})
		return
	} else if err != nil {
		service.Logger.Error("redis测试失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"message": "redis测试失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "redis测试成功", "value": value})
}
