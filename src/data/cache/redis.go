package cache

import (
	"fmt"
	"github.com/melika0-0/bookstore-project/api/config"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) error{
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
	
	DB:       0,  // use default DB
	})
	_, err := redisClient.Ping(redisClient.Context()).Result()
	return err
	
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func CloseRedis() error {
	return redisClient.Close()
}