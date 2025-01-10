package router

import (
	"Adeeb_Go/internal/components/chosen_verses"
	"Adeeb_Go/internal/components/heartbeat"
	"Adeeb_Go/internal/components/poems"
	"Adeeb_Go/internal/components/poets"
	"Adeeb_Go/internal/components/proses"
	"Adeeb_Go/internal/components/users"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

var API huma.API

func InitAPI() huma.API {
	config := huma.DefaultConfig("My API", "0.0.1")
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}
	// // disable the route by setting
	// config.OpenAPIPath = ""

	API = humachi.New(R, config)

	registerAPIs()

	return API
}

func registerAPIs() {
	heartbeat.RegisterAPI(API)
	users.RegisterAPI(API)
	poets.RegisterAPI(API)
	poems.RegisterAPI(API)
	chosen_verses.RegisterAPI(API)
	proses.RegisterAPI(API)
}
