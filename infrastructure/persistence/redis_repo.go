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
