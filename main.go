package main

import (
	"Adeeb_Go/internal/config"
	"Adeeb_Go/internal/database"
	"Adeeb_Go/logger"
	"Adeeb_Go/router"
	"context"
	"fmt"
	"net/http"
)

func main() {
	// Load Config Variables
	config.LoadENV()

	// Init Logger
	logger.Init()

	// Database
	conn, _ := database.Connect()
	defer conn.Close(context.Background())

	// Router & API
	r := router.NewRouter()
	router.UseMiddlewares()
	router.InitAPI()

	logger.Info().Msgf("Starting Server at %v:%v", config.HOST, config.PORT)
	http.ListenAndServe(
		fmt.Sprintf("%v:%v", config.HOST, config.PORT),
		r,
	)
}
