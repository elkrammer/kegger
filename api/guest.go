package api

import (
    "net/http"
    "fmt"
    "strconv"

    "github.com/elkrammer/gorsvp/db"
    "github.com/elkrammer/gorsvp/model"

    "github.com/labstack/echo"
)

// GET - get guests for a party id
func GetGuests(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    db := db.DbManager()

    query := `
    SELECT guests.*
    FROM guests
    where party_refer = $1`

    rows, err := db.Queryx(query, id)

    if err != nil {
        err := fmt.Sprintf("Guest with ID: %v not found", id)
        return echo.NewHTTPError(http.StatusNotFound, err)
    }

    guests := []model.GuestResponse{}
    for rows.Next() {
        g := model.GuestResponse{}
        err = rows.StructScan(&g)
        if err != nil {
            fmt.Println(err)
        }
        guests = append(guests, g)
    }

    return c.JSON(http.StatusOK, guests)
}

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


// POST - create individual guest
func CreateGuest(c echo.Context) error {
    db := db.DbManager()
    guest := new(model.Guest)

    // bind request to model
    if err := c.Bind(guest); err != nil {
        msg := fmt.Sprintf("Invalid request body. %s", err)
        return c.JSON(http.StatusBadRequest, msg)
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

    fmt.Println(guest)

    // insert guest struct into db
    query := `
    INSERT INTO guests
    (first_name, last_name, email, is_attending, party_refer)
    VALUES($1, $2, $3, $4, $5)`
    err := db.QueryRow(query, guest.FirstName, guest.LastName, guest.Email, guest.IsAttending, guest.PartyRefer)

    if err != nil {
        fmt.Println("error inserting guest record: ", query)
        fmt.Println(err)
    }

    return c.JSON(http.StatusOK, guest)
}

// PUT - update guests for a party
func UpdateGuests(c echo.Context) error {
    party_id, _ := strconv.Atoi(c.Param("id"))
    db := db.DbManager()
    guests := []model.Guest{}

    // bind request to model
    if err := c.Bind(&guests); err != nil {
        msg := fmt.Sprintf("Invalid request body. %s", err)
        return c.JSON(http.StatusBadRequest, msg)
    }

    //  update all guests as per the request
    for _, guest := range guests {

        // enforce all guests are under the same party id
        guest.PartyRefer = party_id

        // identify new guests in the request and insert them into the guest table
        if guest.ID == 0 {
            // insert guest struct into db
            query := `
            INSERT INTO guests
            (party_refer, first_name, last_name, email, is_attending)
            VALUES($1, $2, $3, $4, $5)`
            _, err := db.Exec(query, guest.PartyRefer, guest.FirstName, guest.LastName, guest.Email, guest.IsAttending)

            if err != nil {
                fmt.Println("error inserting guest record: ", query)
                fmt.Println(err)
                break
            }
        }

        q := `
        UPDATE guests
        SET first_name=$1, last_name=$2, email=$3, is_attending=$4
        WHERE id=$5;`
        _, err := db.Exec(q, guest.FirstName, guest.LastName, guest.Email, guest.IsAttending, guest.ID)
        if err != nil {
            fmt.Println("error updating guest record: ", q)
            fmt.Println(err)
            break
        }
    }

    return c.JSON(http.StatusOK, guests)
}

// DELETE - delete individual guest
func DeleteGuest(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    db := db.DbManager()

    // check if records exists
    var recordId uint
    query := `SELECT id FROM guests WHERE id = $1`
    err := db.QueryRow(query, id).Scan(&recordId)

    if err != nil && recordId == 0 {
        err := fmt.Sprintf("Guest with ID: %v not found", id)
        return echo.NewHTTPError(http.StatusNotFound, err)
    }

    // delete record
    query = `DELETE FROM guests WHERE id = $1`
    _, err = db.Exec(query, id)

    if err != nil {
        fmt.Println("error deleting guest record: ", query)
        fmt.Println(err)
    }

    return c.JSON(http.StatusOK, H{
        "deleted": id,
    })
}
