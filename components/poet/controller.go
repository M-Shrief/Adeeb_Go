package component_poet

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	not_Available string = "No Poets available"
	not_Found     string = "No Poets found"
)

func getPoets(c echo.Context) error {
	poets := []Poet{}
	err := selectPoets(&poets)
	if err != nil || len(poets) == 0 {
		return c.String(http.StatusNotFound, not_Available)
	}
	return c.JSON(http.StatusOK, poets)
}

func getPoet(c echo.Context) error {
	id := c.Param("id")
	poet := new(Poet)

	err := selectPoet(c.Request().Context(), id, poet)
	if err != nil {
		return c.String(http.StatusNotFound, not_Found)
	}
	return c.JSON(http.StatusOK, poet)
}
