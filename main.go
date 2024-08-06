package main

import (
	"golangredis/domain/model"
	"golangredis/infrastructure/persistence"
	"golangredis/interfaces"
	"golangredis/usecase"
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

	//load config
	err := configor.Load(&cfg, "config.yml")
	if err != nil {
		log.Println("error config ", err)
		return
	}

	//get repositories
	newRepo, err := persistence.NewRepositories()
	if err != nil {
		log.Println("redis init error", err)
	}

	//new usecase
	redisUsecase := usecase.NewRedisUsecase(newRepo.RedisRepo)

	//new handler
	redisHandler := interfaces.NewRedisHandler(redisUsecase)

	//config fiber
	fiberCfg := fiber.Config{
		BodyLimit: 5 * 1024 * 1024,
	}

	//new fiber instance
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

	app.Post("redis/set", redisHandler.SetHandler)

	app.Listen(":" + cfg.Port)
}
