package api

import (
    "net/http"
    "fmt"
    "encoding/json"
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
    parties := []model.PartyResponse{}

    query := `
    SELECT
    parties.*,
    users.name as host_name,
    (SELECT json_agg(row_to_json(guests))
    FROM guests
    WHERE guests.party_refer = parties.id) AS guests
    FROM parties
    INNER JOIN users ON parties.host_id = users.id`

    rows, err := db.Queryx(query)

    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "No parties found")
    }

    for rows.Next() {
        err = rows.StructScan(&party)
        if err != nil {
            fmt.Println(err)
        }

        guest := []model.Guest{}
        err := json.Unmarshal([]byte(party.Guests), &guest)
        if err != nil {
            fmt.Println(err)
        }

        // TODO: this is fucking ugly. find out a better way to do it.
        p := model.PartyResponse{}
        p.ID = party.ID
        p.Name = party.Name
        p.InvitationId = party.InvitationId
        p.InvitationOpened = party.InvitationOpened
        p.IsAttending = party.IsAttending
        p.HostName = party.HostName
        p.Comments = party.Comments
        p.Guests = guest

        parties = append(parties, p)

    }

    return c.JSON(http.StatusOK, parties)
}

// GET - get an individual party
func GetParty(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    db := db.DbManager()
    party := model.Party{}

    query := `
    SELECT
    parties.*,
    users.name as host_name,
    (SELECT json_agg(row_to_json(guests))
    FROM guests
    WHERE guests.party_refer = parties.id) AS guests
    FROM parties
    INNER JOIN users ON parties.host_id = users.id
    WHERE parties.id = $1`

    rows, err := db.Queryx(query, id)

    if err != nil {
        err := fmt.Sprintf("Party with ID: %v not found", id)
        return echo.NewHTTPError(http.StatusNotFound, err)
    }

    p := model.PartyResponse{}
    for rows.Next() {
        err = rows.StructScan(&party)
        if err != nil {
            fmt.Println(err)
        }

        guest := []model.Guest{}
        err := json.Unmarshal([]byte(party.Guests), &guest)
        if err != nil {
            fmt.Println(err)
        }

        // TODO: this is kinda ugly. find out a better way to do it.
        p.ID = party.ID
        p.Name = party.Name
        p.InvitationId = party.InvitationId
        p.InvitationOpened = party.InvitationOpened
        p.IsAttending = party.IsAttending
        p.HostName = party.HostName
        p.Comments = party.Comments
        p.Guests = guest
    }

    return c.JSON(http.StatusOK, p)
}

// POST - create new party
func CreateParty(c echo.Context) error {
    party := model.PartyRequest{}
    db := db.DbManager()
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

    // insert party struct into db
    // TODO: figure out a more elegant way to do this
    var pid uint
    query := `
    INSERT INTO parties
    ("name", invitation_id, invitation_sent, invitation_opened, is_attending, "comments", host_id)
    VALUES($1, $2, $3, $4, $5, $6, $7)
    RETURNING id`
    err := db.QueryRow(query, party.Name, party.InvitationId, party.InvitationSent, party.InvitationOpened, party.IsAttending, party.Comments, party.HostId).Scan(&pid)

    if err != nil {
        fmt.Println("error inserting party record: ", query)
        fmt.Println(err)
    }

    // insert all guests in the request into the guests table
    for _, guest := range party.Guests {
        q := `
        INSERT INTO guests
        (party_refer, first_name, last_name, email, is_attending)
        VALUES($1, $2, $3, $4, $5);`
        db.Exec(q, pid, guest.FirstName, guest.LastName, guest.Email, guest.IsAttending)
    }

    return c.JSON(http.StatusCreated, party)
}

/*
// PUT - update party
func UpdateParty(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    db := db.DbManager()

    party := new(model.Party)

    // check if records exists
    if db.First(&party, id).RecordNotFound() {
        err := fmt.Sprintf("Party with ID: %v not found", id)
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
        err := fmt.Sprintf("Party with ID: %v not found", id)
        return c.JSON(http.StatusBadRequest, err)
    }

    db.Delete(&party)

    return c.JSON(http.StatusOK, H{
        "deleted": id,
    })

}

*/
