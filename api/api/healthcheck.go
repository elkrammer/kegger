package api

import (
	"log"
	"net/http"

	"github.com/elkrammer/kegger/api/db"

	"github.com/labstack/echo/v4"
)

// GET - Health Check
func HealthCheck(c echo.Context) error {
	db := db.DbManager()
	err := db.Ping()

	if err != nil {
		log.Print(err)
		return c.String(http.StatusBadRequest, "Kegger API is down :(")
	}

	return c.String(http.StatusOK, "Kegger API is up! :)")
}
