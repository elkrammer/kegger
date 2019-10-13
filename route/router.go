package route

import (
    "github.com/elkrammer/gorsvp/api"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routes
	e.GET("/", api.GetParties)
	e.POST("/", api.CreateParty)

	return e
}
