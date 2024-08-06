package repository

import (
	"context"
	"golangredis/domain/model"
)

type IRedisRepository interface {
	SetValue(ctx context.Context, req model.SetValueReq) error
	GetValue(ctx context.Context, key string) (string, error)
}
