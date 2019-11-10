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

        // TODO: this is fucking ugly. find a better way to do it.
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
        msg := fmt.Sprintf("Invalid request body. %s", err)
        return c.JSON(http.StatusBadRequest, msg)
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
        _, err = db.Exec(q, pid, guest.FirstName, guest.LastName, guest.Email, guest.IsAttending)
        if err != nil {
            fmt.Println("error inserting guest record: ", q)
            fmt.Println(err)
        }
    }

    // return a PartyResponse with the newly created party
    query = `
    SELECT
    parties.*,
    users.name as host_name,
    (SELECT json_agg(row_to_json(guests))
    FROM guests
    WHERE guests.party_refer = parties.id) AS guests
    FROM parties
    INNER JOIN users ON parties.host_id = users.id
    WHERE parties.id = $1`

    rows, err := db.Queryx(query, pid)

    if err != nil {
        err := fmt.Sprintf("Party with ID: %v not found", pid)
        return echo.NewHTTPError(http.StatusNotFound, err)
    }

    p := model.Party{}
    r := model.PartyResponse{}
    for rows.Next() {
        err = rows.StructScan(&p)
        if err != nil {
            fmt.Println(err)
        }

        guest := []model.Guest{}
        err := json.Unmarshal([]byte(p.Guests), &guest)
        if err != nil {
            fmt.Println(err)
        }

        // TODO: this is kinda ugly. find a better way to do it.
        r.ID = p.ID
        r.Name = p.Name
        r.InvitationId = p.InvitationId
        r.InvitationOpened = p.InvitationOpened
        r.IsAttending = p.IsAttending
        r.HostName = p.HostName
        r.Comments = p.Comments
        r.Guests = guest
    }

    return c.JSON(http.StatusCreated, r)
}

// PUT - update party
func UpdateParty(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    db := db.DbManager()

    party := new(model.PartyRequest)

    // check if records exists
    var recordId uint
    query := `SELECT id FROM parties WHERE id = $1`
    err := db.QueryRow(query, id).Scan(&recordId)

    if err != nil && recordId == 0 {
        err := fmt.Sprintf("Party with ID: %v not found", id)
        return echo.NewHTTPError(http.StatusNotFound, err)
    }

    // bind request to model
    if err := c.Bind(party); err != nil {
        return c.JSON(http.StatusBadRequest, party)
    }

    query = `
    UPDATE parties
    SET "name"=$1, invitation_id=$2, invitation_sent=$3, invitation_opened=$4, is_attending=$5, "comments"=$6, host_id=$7
    WHERE id=$8;`
    _, err = db.Exec(query, party.Name, party.InvitationId, party.InvitationSent, party.InvitationOpened, party.IsAttending, party.Comments, party.HostId, id)

    if err != nil {
        fmt.Println("error updating party record: ", query)
        fmt.Println(err)
    }

    //  update all guests as per the request
    for _, guest := range party.Guests {
        q := `
        UPDATE guests
        SET party_refer=$1, first_name=$2, last_name=$3, email=$4, is_attending=$5
        WHERE id=$6;`
        _, err := db.Exec(q, id, guest.FirstName, guest.LastName, guest.Email, guest.IsAttending, guest.ID)
        if err != nil {
            fmt.Println("error updating guest record: ", q)
            fmt.Println(err)
        }
    }

    return c.JSON(http.StatusOK, party)
}

// DELETE - delete party
func DeleteParty(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    db := db.DbManager()

    // check if records exists
    var recordId uint
    query := `SELECT id FROM parties WHERE id = $1`
    err := db.QueryRow(query, id).Scan(&recordId)

    if err != nil && recordId == 0 {
        err := fmt.Sprintf("Party with ID: %v not found", id)
        return echo.NewHTTPError(http.StatusNotFound, err)
    }

    // delete record
    query = `DELETE FROM parties WHERE id = $1`
    _, err = db.Exec(query, id)

    if err != nil {
        fmt.Println("error deleting party record: ", query)
        fmt.Println(err)
    }

    return c.JSON(http.StatusOK, H{
        "deleted": id,
    })
}
