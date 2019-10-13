package api

import (
    "net/http"
    // "fmt"

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

func CreateParty(c echo.Context) error {
    db := db.DbManager()
    party := model.Party{}
    if err := c.Bind(&party); err != nil {
        return err
    }

    if party.Name == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "Missing field Name")
    }

    if len(party.Guests) == 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Missing field Guests. You must include at least one Guest")
    }

    db.Create(&party)

    return c.JSON(http.StatusCreated, party)
}

