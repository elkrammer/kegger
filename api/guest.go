package api

import (
    "net/http"
    "fmt"
    "strconv"

    "github.com/elkrammer/gorsvp/db"
    "github.com/elkrammer/gorsvp/model"

    "github.com/labstack/echo"
)

// GET - get individual guest
func GetGuest(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    db := db.DbManager()
    guest := model.Guest{}

    if db.First(&guest, id).RecordNotFound() {
        err := fmt.Sprintf("Guest with ID: %v not found", id)
        return echo.NewHTTPError(http.StatusNotFound, err)
    }

    if err := db.First(&guest, id).Error; err != nil {
        return err
    }

    return c.JSON(http.StatusOK, guest)
}
