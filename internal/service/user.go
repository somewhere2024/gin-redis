package service

import (
	"gin-redis/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var jwtkey = []byte("MySecret")
var Logger *zap.Logger

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	}
}

// 验证密码
func VerifyPassword(password string, User *models.User) bool {
	if password == User.Password {
		return true
	} else {
		return false
	}
}
func CreateToken(data jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString(jwtkey)
}

func ParseTokne(token string) (jwt.MapClaims, error) {
	user := jwt.MapClaims{}
	jwtstr, err := jwt.ParseWithClaims(token, user, secret())
	if err != nil {
		Logger.Error("Error parsing token", zap.Error(err))
		return nil, err
	}
	if !jwtstr.Valid {
		Logger.Error("Invalid token")
		return nil, err
	}
	return user, nil
}

func LogInit() {
	Logger, _ = zap.NewProduction()
	Logger.Info("Starting server")
}

// 生成UUID用户唯一ID
func CreateUUId() string {
	id := uuid.New()
	return id.String()
}
