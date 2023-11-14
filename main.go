package main

import (
	component_poet "my-way/components/poet"
	"my-way/config"
	"my-way/datasource"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Could not load environment variables config.")
	}

	config.AppConfig = &conf

	datasource.ConnectDB()
	defer datasource.DB.Close()

	datasource.ConnectRedis()

	app := fiber.New()

	app.Use(helmet.New())
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}]: ${method} ${path} - Status: ${status}, Latency: ${latency}. \n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	api := app.Group("/api")
	component_poet.Init(api)

	err = app.Listen(":1323")
	if err != nil {
		log.Fatal().Msgf("Can't initialized app, error:%v", err)
	}
}
