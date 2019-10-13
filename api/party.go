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

    parties := []model.Party{}
    db.Find(&parties)

    return c.JSON(http.StatusOK, parties)
}
