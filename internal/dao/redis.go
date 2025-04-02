package dao

import "github.com/go-redis/redis/v8"

var RDB *redis.Client

func InitRedisDB() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "124.70.204.2:6379",
		Password: "",
		DB:       0,
	})
}
