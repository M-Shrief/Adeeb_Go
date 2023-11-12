package component_poet

import "github.com/labstack/echo/v4"

func Init(e *echo.Echo) {
	e.GET("/poet/:id", getPoet)
	e.GET("/poets", getPoets)
}
