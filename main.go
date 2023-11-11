package main

import (
	"log"
	component_poet "my-way/components/poet"
	"my-way/config"
	"my-way/datasource"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {

	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables config.", err)
	}

	config.AppConfig = &conf

	datasource.ConnectDB()
	defer datasource.DB.Close()

	datasource.ConnectRedis()

	e := echo.New()

	e.Use(middleware.Secure())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "METHOD=${method}, URI=${uri}, STATUS=${status}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	component_poet.Init(e)

	e.Logger.Fatal(e.Start(":1323"))
}
