package usecase

import (
	"context"
	"errors"
	"golangredis/domain/model"
	"golangredis/domain/repository"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type redisUsecase struct {
	redisRepo repository.IRedisRepository
}

type IRedisUsecase interface {
	SetUsecase(ctx context.Context, req model.RedisReq) *model.BaseResponse
}

func NewRedisUsecase(redisRepo repository.IRedisRepository) IRedisUsecase {
	return &redisUsecase{
		redisRepo: redisRepo,
	}
}

// SetUsecase implements IRedisUsecase
func (u *redisUsecase) SetUsecase(ctx context.Context, req model.RedisReq) *model.BaseResponse {
	usecaseName := "[SetUsecase]"

	expireDuration, err := time.ParseDuration(req.ExpireTime)
	if err != nil {
		log.Error(usecaseName+" SetValue error: ", err)

		return new(model.BaseResponse).Response(err, nil)
	}

	setValueReq := model.SetValueReq{
		Key:        req.Key,
		Value:      req.Value,
		ExpireTime: expireDuration,
	}

	err = u.redisRepo.SetValue(ctx, setValueReq)
	if err != nil {
		log.Error(usecaseName+" SetValue error: ", err)

		return new(model.BaseResponse).Response(errors.New("Failed set redis"), nil)
	}

	return new(model.BaseResponse).Response(nil, "Success")
}
