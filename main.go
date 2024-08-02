package main

import (
	"golangredis/domain/model"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"github.com/jinzhu/configor"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	cfg := model.RedisConfig{}

	err := configor.Load(&cfg, "config.yml")
	if err != nil {
		log.Println("error config ", err)
		return
	}

	fiberCfg := fiber.Config{
		BodyLimit: 5 * 1024 * 1024,
	}

	app := fiber.New(fiberCfg)

	app.Use(logger.New())
	app.Use(helmet.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:6001",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		ExposeHeaders:    "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Header",
		AllowCredentials: true,
		MaxAge:           86400,
	}))

	app.Listen(":" + cfg.Port)
}
