package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/elkrammer/gorsvp/db"
	"github.com/elkrammer/gorsvp/model"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
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
		p.HostId = party.HostId
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
		p.HostId = party.HostId
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
    ("name", host_id, "comments")
    VALUES($1, $2, $3)
    RETURNING id`
	err := db.QueryRow(query, party.Name, party.HostId, party.Comments).Scan(&pid)

	if err != nil {
		fmt.Println("error inserting party record: ", query)
		fmt.Println(err)
	}

	// insert all guests in the request into the guests table
	for _, guest := range party.Guests {
		invitationId := ksuid.New()
		q := `
        INSERT INTO guests
        (party_refer, first_name, last_name, email, is_attending, invitation_id, plus_one)
        VALUES($1, $2, $3, $4, $5, $6, $7);`
		_, err = db.Exec(q, pid, guest.FirstName, guest.LastName, guest.Email, guest.IsAttending, invitationId, guest.PlusOne)
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
    SET "name"=$1, "comments"=$2, host_id=$3
    WHERE id=$4;`
	_, err = db.Exec(query, party.Name, party.Comments, party.HostId, id)

	if err != nil {
		fmt.Println("error updating party record: ", query)
		fmt.Println(err)
	}

	//  update all guests as per the request
	for _, guest := range party.Guests {
		q := `
        UPDATE guests
        SET party_refer=$1, first_name=$2, last_name=$3, email=$4, is_attending=$5, plus_one=$6
        WHERE id=$7;`
		_, err := db.Exec(q, id, guest.FirstName, guest.LastName, guest.Email, guest.IsAttending, guest.PlusOne, guest.ID)
		if err != nil {
			fmt.Println("error updating guest record: ", q)
			fmt.Println(err)
		}
	}

	// return a PartyResponse
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

	rows, err := db.Queryx(query, recordId)

	if err != nil {
		err := fmt.Sprintf("Party with ID: %v not found", recordId)
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
		r.HostName = p.HostName
		r.Comments = p.Comments
		r.Guests = guest
	}

	return c.JSON(http.StatusOK, r)
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
