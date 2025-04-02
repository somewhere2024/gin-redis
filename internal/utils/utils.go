package utils

import (
	"context"
	"fmt"
	"gin-redis/internal/service"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// 存储键值
func SetKeyValue(rdb *redis.Client, key string, value string) error {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// 获取键值
func GetKeyValue(rdb *redis.Client, key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		service.Logger.Info("key does not exist")
		return "", nil
	} else if err != nil {
		return "", err
	}
	fmt.Println(val)
	return val, nil
}
