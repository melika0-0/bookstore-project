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
	Password: cfg.Redis.Password,
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
package cache

import (
	"fmt"
	"time"
	"github.com/melika0-0/bookstore-project/api/config"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) error{
	redisClient = redis.NewClient(&redis.Options{
	Addr:  fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port) ,
	Password: cfg.Redis.Password,
	DB:       0,  // use default DB
	DialTimeout: cfg.Redis.dialTimeout * time.Second,
	ReadTimeout: cfg.Redis.readTimeout * time.Second,
	WriteTimeout: cfg.Redis.writeTimeout * time.Second,
	IdleCheckFrequency:   cfg.Redis.idlecheckfrequency * time.Millisecond,
	PoolSize: cfg.Redis.poolsize,
	PoolTimeout: cfg.Redis.pooltimeout,
	MinIdleConns: cfg.Redis.minIdleConnections,
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