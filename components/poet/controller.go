package component_poet

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func getPoets(c echo.Context) error {
	poets := []Poet{}
	err := selectPoets(&poets)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusOK, "No Poets found")
	}
	return c.JSON(http.StatusOK, poets)
}

func getPoet(c echo.Context) error {
	id := c.Param("id")
	poet := new(Poet)

	err := selectPoet(c.Request().Context(), id, poet)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusOK, "Can't Find Poet")
	}
	return c.JSON(http.StatusOK, poet)
}
