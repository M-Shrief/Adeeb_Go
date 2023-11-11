package main

import (
	"log"
	component_poet "my-way/components/poet"
	"my-way/config"
	"my-way/datasource"
	"net/http"

	"github.com/labstack/echo"
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
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	component_poet.Init(e)

	e.Logger.Fatal(e.Start(":1323"))
}
