package api

import (
	"net/http"

	"github.com/elkrammer/kegger/api/db"

	"github.com/labstack/echo/v4"
)

// GET - Health Check
func HealthCheck(c echo.Context) error {
	db := db.DbManager()
	query := `SELECT * from settings;`
	_, err := db.Queryx(query)

	if err != nil {
		return c.String(http.StatusBadRequest, "Kegger API is DOWN")
	}

	return c.String(http.StatusOK, "Kegger API is UP")
}
