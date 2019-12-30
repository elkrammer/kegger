package route

import (
	"os"

	"github.com/elkrammer/gorsvp/api"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jwtClaim struct {
	Username string `json:"username"`
	UserID   uint   `json:"sub"`
	jwt.StandardClaims
}

func Init() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization, echo.HeaderAccept},
		AllowCredentials: true,
		AllowMethods:     []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
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
	r.POST("/guest", api.CreateGuest)
	r.GET("/guests/:id", api.GetGuests)
	r.GET("/guest/:id", api.GetGuest)
	r.PUT("/guest/:id", api.UpdateGuests)
	r.DELETE("/guest/:id", api.DeleteGuest)

	// User Routes
	r.POST("/user", api.CreateUser)
	r.GET("/user", api.GetUsers)
	r.GET("/user/:id", api.GetUser)
	r.PUT("/user/:id", api.EditUser)
	r.DELETE("/user/:id", api.DeleteUser)

	// Settings Routes
	r.GET("/settings", api.GetSettings)
	r.PUT("/settings", api.UpdateSettings)

	// Send Invite Email Route
	r.POST("/sendinvite/:id", api.SendInvite)

	// Fetch Invitation Details
	e.GET("/invite/:invite_id", api.FetchInviteDetails)

	return e
}
