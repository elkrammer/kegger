package api

import (
    "net/http"
    "fmt"

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
