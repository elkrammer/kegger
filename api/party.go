package api

import (
    "net/http"
    "fmt"
    "strconv"

    "github.com/elkrammer/gorsvp/db"
    "github.com/elkrammer/gorsvp/model"

    "github.com/labstack/echo"
)

type H map[string]interface{}

// GET - get all parties
func GetParties(c echo.Context) error {
    db := db.DbManager()
    party := model.Party{}
    parties := []model.Party{}

    if err := db.Preload("Guests").Find(&parties).Error; err != nil {
        return err
    }

    if db.Find(&party).RecordNotFound() {
        return echo.NewHTTPError(http.StatusNotFound, "No parties found")
    }

    return c.JSON(http.StatusOK, parties)
}

// GET - get individual party
func GetParty(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    db := db.DbManager()
    party := model.Party{}

    if db.First(&party, id).RecordNotFound() {
        err := fmt.Sprintf("Party with ID: %s not found", id)
        return echo.NewHTTPError(http.StatusNotFound, err)
    }

    if err := db.Preload("Guests").First(&party, id).Error; err != nil {
        return err
    }

    return c.JSON(http.StatusOK, party)
}

// POST - create new party
func CreateParty(c echo.Context) error {
    db := db.DbManager()
    party := model.Party{}
    if err := c.Bind(&party); err != nil {
        return err
    }

    // request validations
    if party.Name == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "Missing field Name")
    }

    if len(party.Guests) == 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Missing field Guests. You must include at least one Guest")
    }

    db.Create(&party)

    return c.JSON(http.StatusCreated, party)
}

// PUT - update party
func UpdateParty(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    db := db.DbManager()

    party := new(model.Party)

    // check if records exists
    if db.First(&party, id).RecordNotFound() {
        err := fmt.Sprintf("Party with ID: %s not found", id)
        return echo.NewHTTPError(http.StatusNotFound, err)
    }

    // check for errors
    if err := db.First(&party, id).Error; err != nil {
        return err
    }

    // bind request to model
    if err := c.Bind(party); err != nil {
        return c.JSON(http.StatusBadRequest, party)
    }

    db.Save(&party)
    return c.JSON(http.StatusOK, party)
}

// DELETE - delete party
func DeleteParty(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    db := db.DbManager()
    party := new(model.Party)

    // check if record exists
    if db.First(&party, id).RecordNotFound() {
        err := fmt.Sprintf("Party with ID: %s not found", id)
        return c.JSON(http.StatusBadRequest, err)
    }

    db.Delete(&party)

    return c.JSON(http.StatusOK, H{
        "deleted": id,
    })

}
