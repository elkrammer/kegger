package route

import (
	"os"

	"github.com/elkrammer/kegger/api/api"

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

	// Invitation Routes
	r.POST("/sendinvite/:id", api.SendInvite)           // Send Invite Email Route (protected)
	e.GET("/invite/:invite_id", api.FetchInviteDetails) // Fetch Invitation Details
	e.PUT("/invite", api.UpdateInvite)                  // Update Invitation Attendance / Invite Opened Timestamp
	e.GET("/findinvite/:email", api.FindInviteId)       // Find Invitation ID from email

	// Health Check
	e.GET("/health", api.HealthCheck)

	return e
}
