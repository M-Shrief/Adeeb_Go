package component_poet

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	not_Available string = "No Poets available"
	not_Found     string = "No Poets found"
)

func getPoets(c *fiber.Ctx) error {
	poets := []Poet{}
	err := selectPoets(&poets)
	if err != nil || len(poets) == 0 {
		return c.Status(http.StatusNotFound).SendString(not_Available)
	}
	return c.Status(http.StatusOK).JSON(poets)
}

func getPoet(c *fiber.Ctx) error {
	id := c.Params("id")
	poet := new(Poet)

	err := selectPoet(c.Context(), id, poet)
	if err != nil {
		return c.Status(http.StatusNotFound).SendString(not_Found)
	}
	return c.Status(http.StatusOK).JSON(poet)
}
