package model

import (
	"github.com/gofiber/fiber/v2"
)

type RedisConfig struct {
	Host string `env:"host"`
	Port string `env:"port"`
}

type BaseResponse struct {
	ErrorMsg string      `json:"error"`
	Data     interface{} `json:"data"`
}

func (br *BaseResponse) Response(err error, data interface{}) *BaseResponse {
	return &BaseResponse{ErrorMsg: err.Error(), Data: data}
}

func Response(ctx *fiber.Ctx, statusHttp int, data interface{}) error {
	return ctx.Status(statusHttp).JSON(data)
}
