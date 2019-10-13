package api

import (
    "net/http"
    //"fmt"

    "github.com/elkrammer/gorsvp/db"
    "github.com/elkrammer/gorsvp/model"

    "github.com/labstack/echo"
)

func GetParties(c echo.Context) error {
    db := db.DbManager()
    party := model.Party{}
    parties := []model.Party{}

    if err := db.Find(&parties).Error; err != nil {
        return err
    }

    if db.Find(&party).RecordNotFound() {
        return echo.NewHTTPError(http.StatusNotFound, "No parties found")
    }

    return c.JSON(http.StatusOK, parties)
}
