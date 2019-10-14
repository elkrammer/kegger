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

    // Party Routes
    e.GET("/party", api.GetParties)
    e.POST("/party", api.CreateParty)
    e.GET("/party/:id", api.GetParty)
    e.PUT("/party/:id", api.UpdateParty)
    e.DELETE("/party/:id", api.DeleteParty)

    // Guest Routes
    e.GET("/guest/:id", api.GetGuest)
    e.POST("/guest", api.CreateGuest)
    e.PUT("/guest/:id", api.UpdateGuest)
    e.DELETE("/guest/:id", api.DeleteGuest)

    return e
}
