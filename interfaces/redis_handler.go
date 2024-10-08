package interfaces

import (
	"golangredis/domain/model"
	"golangredis/infrastructure/middlewares"
	"golangredis/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
)

type redisHandler struct {
	redisUc usecase.IRedisUsecase
}

func NewRedisHandler(redisUc usecase.IRedisUsecase) *redisHandler {
	return &redisHandler{
		redisUc: redisUc,
	}
}

func (h *redisHandler) SetHandler(ctx *fiber.Ctx) error {
	handlerName := "[SetHandler]"

	req := model.RedisReq{}
	if err := ctx.BodyParser(&req); err != nil {
		log.Error(handlerName+" BodyParser error: ", err)

		return model.Response(ctx, http.StatusBadRequest, err)
	}

	if err := middlewares.ValidateRequest(req); err != nil {
		log.Error(handlerName+" ValidateRequest error: ", err)

		return model.Response(ctx, http.StatusBadRequest, err)
	}

	return model.Response(ctx, http.StatusOK, h.redisUc.SetUsecase(ctx.Context(), req))
}

func (h *redisHandler) GetHandler(ctx *fiber.Ctx) error {
	handlerName := "[GetHandler]"

	req := model.RedisReq{}
	if err := ctx.BodyParser(&req); err != nil {
		log.Error(handlerName+" BodyParser error: ", err)

		return model.Response(ctx, http.StatusBadRequest, err)
	}

	if err := middlewares.ValidateRequest(req); err != nil {
		log.Error(handlerName+" ValidateRequest error: ", err)

		return model.Response(ctx, http.StatusBadRequest, err)
	}

	return model.Response(ctx, http.StatusOK, h.redisUc.GetUsecase(ctx.Context(), req.Key))
}

func (h *redisHandler) DeleteHandler(ctx *fiber.Ctx) error {
	handlerName := "[DeleteHandler]"

	req := model.RedisReq{}
	if err := ctx.BodyParser(&req); err != nil {
		log.Error(handlerName+" BodyParser error: ", err)

		return model.Response(ctx, http.StatusBadRequest, err)
	}

	if err := middlewares.ValidateRequest(req); err != nil {
		log.Error(handlerName+" ValidateRequest error: ", err)

		return model.Response(ctx, http.StatusBadRequest, err)
	}

	return model.Response(ctx, http.StatusOK, h.redisUc.DeleteUsecase(ctx.Context(), req.Key))
}
