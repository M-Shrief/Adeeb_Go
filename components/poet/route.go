package component_poet

import "github.com/labstack/echo"

func Init(e *echo.Echo) {
	e.GET("/poet/:id", getPoet)
	e.GET("/poets", getPoets)
}
