package main

import (
	component_poet "my-way/components/poet"
	"my-way/config"
	"my-way/datasource"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e := echo.New()

	e.Use(middleware.Secure())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("URI", v.URI).
				Str("latency", v.Latency.String()).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	component_poet.Init(e)

	e.Logger.Fatal(e.Start(":1323"))
}
