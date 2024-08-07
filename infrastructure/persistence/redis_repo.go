package persistence

import (
	"context"
	"golangredis/domain/model"
	"golangredis/domain/repository"

	redis "github.com/redis/go-redis/v9"
)

type redisRepository struct {
	redisClient *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) repository.IRedisRepository {
	return &redisRepository{
		redisClient: redisClient,
	}
}

// SetValue implements repository.IRedisRepository
func (r *redisRepository) SetValue(ctx context.Context, req model.SetValueReq) error {
	return r.redisClient.Set(ctx, req.Key, req.Value, req.ExpireTime).Err()
}

func (r *redisRepository) GetValue(ctx context.Context, key string) (string, error) {
	redisValue, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return redisValue, nil
}

func (r *redisRepository) DeleteValue(ctx context.Context, key string) error {
	return r.redisClient.Del(ctx, key).Err()
}
