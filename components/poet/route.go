package component_poet

import "github.com/gofiber/fiber/v2"

func Init(api fiber.Router) {
	api.Get("/poets", getPoets)
	api.Get("/poet/:id", getPoet)
}
