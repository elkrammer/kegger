package api

import (
	"fmt"
	"net/http"

	"github.com/elkrammer/gorsvp/db"
	"github.com/elkrammer/gorsvp/model"

	"github.com/labstack/echo"
)

// GET - get all settings and values
func GetSettings(c echo.Context) error {
	db := db.DbManager()

	query := `SELECT * from settings;`
	rows, err := db.Queryx(query)

	if err != nil {
		err := fmt.Sprintf("No settings found")
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	settings := []model.Settings{}
	for rows.Next() {
		s := model.Settings{}
		err = rows.StructScan(&s)
		if err != nil {
			fmt.Println(err)
		}
		settings = append(settings, s)
	}

	return c.JSON(http.StatusOK, settings)
}

func UpdateSettings(c echo.Context) error {
	db := db.DbManager()
	settings := []model.Settings{}

	// bind request to model
	if err := c.Bind(&settings); err != nil {
		msg := fmt.Sprintf("Invalid request body. %s", err)
		return c.JSON(http.StatusBadRequest, msg)
	}

	//  update all keys as per the request
	for _, setting := range settings {
		query := `UPDATE settings SET "name" = $1, "value" = $2 WHERE "name" = $3`
		_, err := db.Exec(query, setting.Name, setting.Value, setting.Name)
		if err != nil {
			fmt.Println("error updating setting: ", query)
			fmt.Println(err)
			break
		}
	}

	return c.JSON(http.StatusOK, settings)
}
