package persistence

import (
	"context"
	"golangredis/domain/repository"

	"github.com/redis/go-redis/v9"
)

type Repositories struct {
	redisClient *redis.Client
	RedisRepo   repository.IRedisRepository
}

func NewRepositories() (*Repositories, error) {
	newRedis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &Repositories{redisClient: newRedis, RedisRepo: NewRedisRepository(newRedis)}, newRedis.Ping(context.Background()).Err()
}
