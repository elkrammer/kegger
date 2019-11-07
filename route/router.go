package route

import (
    "os"

    "github.com/elkrammer/gorsvp/api"

    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

type jwtClaim struct {
    Username  string `json:"username"`
    UserID    uint `json:"sub"`
    jwt.StandardClaims
}

func Init() *echo.Echo {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization, echo.HeaderAccept},
        AllowCredentials: true,
        AllowMethods: []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
    }))

    // Authentication
    e.POST("/login", api.UserLogin)
    e.POST("/token", api.GenRefreshToken)

    // Restricted routes
    r := e.Group("/api")
    config := middleware.JWTConfig{
        Claims:     &jwtClaim{},
        SigningKey: []byte(os.Getenv("JWT_SECRET")),
    }
    r.Use(middleware.JWTWithConfig(config))

    // Party Routes
    r.POST("/party", api.CreateParty)
    r.GET("/party", api.GetParties)
    r.GET("/party/:id", api.GetParty)
    r.PUT("/party/:id", api.UpdateParty)
    r.DELETE("/party/:id", api.DeleteParty)

    // Guest Routes
    r.GET("/guest/:id", api.GetGuest)
    /*
    r.POST("/guest", api.CreateGuest)
    r.PUT("/guest/:id", api.UpdateGuest)
    r.DELETE("/guest/:id", api.DeleteGuest)
    */
    return e
}
