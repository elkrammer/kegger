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

    query := `
    SELECT guests.*, parties.name as party_name
    FROM guests
    INNER JOIN parties ON guests.party_refer = parties.id
    where parties.id = $1`

    rows, err := db.Queryx(query, id)

    if err != nil {
        err := fmt.Sprintf("Guest with ID: %v not found", id)
        return echo.NewHTTPError(http.StatusNotFound, err)
    }

    g := model.GuestResponse{}
    for rows.Next() {
        err = rows.StructScan(&g)
        if err != nil {
            fmt.Println(err)
        }
    }

    return c.JSON(http.StatusOK, g)
}

/*

// POST - create individual guest
func CreateGuest(c echo.Context) error {
    db := db.DbManager()
    guest := new(model.Guest)

    // bind request to model
    if err := c.Bind(guest); err != nil {
        return c.JSON(http.StatusBadRequest, guest)
    }

    // request validations
    if guest.FirstName == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "Missing FirstName field")
    }

    if guest.LastName == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "Missing LastName field")
    }

    if guest.Email == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "Missing Email field")
    }

    if guest.PartyRefer == 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Missing PartyRefer field containing party id")
    }

    db.Create(&guest)

    return c.JSON(http.StatusOK, guest)
}


// PUT - update individual guest
func UpdateGuest(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    db := db.DbManager()
    guest := new(model.Guest)

    if db.First(&guest, id).RecordNotFound() {
        err := fmt.Sprintf("Guest with ID: %v not found", id)
        return echo.NewHTTPError(http.StatusNotFound, err)
    }

    // check for errors
    if err := db.First(&guest, id).Error; err != nil {
        return err
    }

    // bind request to model
    if err := c.Bind(guest); err != nil {
        return c.JSON(http.StatusBadRequest, guest)
    }

    db.Save(&guest)

    return c.JSON(http.StatusOK, guest)
}

// DELETE - delete individual guest
func DeleteGuest(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    db := db.DbManager()
    guest := new(model.Guest)

    // check if record exists
    if db.First(&guest, id).RecordNotFound() {
        err := fmt.Sprintf("Party with ID: %v not found", id)
        return c.JSON(http.StatusBadRequest, err)
    }

    db.Delete(&guest)

    return c.JSON(http.StatusOK, H{
        "deleted": id,
    })
}

*/

