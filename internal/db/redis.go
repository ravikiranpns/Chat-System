package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitRedis(host, port string) error {
	opt := &redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	}
	redisClient = redis.NewClient(opt)

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}
