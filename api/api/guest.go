package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/elkrammer/kegger/api/db"
	"github.com/elkrammer/kegger/api/internal/helper"
	"github.com/elkrammer/kegger/api/model"

	"github.com/labstack/echo/v4"
	"github.com/segmentio/ksuid"
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
	where guests.id = $1`

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

	// insert guest struct into db
	invitationId := ksuid.New()
	query := `
	INSERT INTO guests
	(first_name, last_name, email, is_attending, party_refer, invitation_id, plus_one)
	VALUES($1, $2, $3, $4, $5, $6, $7)`
	err := db.QueryRow(query, guest.FirstName, guest.LastName, guest.Email, guest.IsAttending, guest.PartyRefer, invitationId, guest.PlusOne)

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

	// first let's get a list of all the ids of the guests that currently belong to this party
	var currentGuests []int
	query := `
	SELECT id
	FROM guests
	where party_refer = $1`
	rows, err := db.Queryx(query, party_id)

	if err != nil {
		err := fmt.Sprintf("no guests for party %v found", party_id)
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	for rows.Next() {
		var id int
		rows.Scan(&id)
		currentGuests = append(currentGuests, id)
	}

	var requestGuests []int // keep track of the guest ids in the request
	//  update all guests as per the request
	for _, guest := range guests {
		requestGuests = append(requestGuests, guest.ID)
		guest.PartyRefer = party_id // enforce all guests are under the same party id

		// identify new guests in the request and insert them into the guest table
		if guest.ID == 0 {
			invitationId := ksuid.New()
			// insert guest struct into db
			query := `
					INSERT INTO guests
					(party_refer, first_name, last_name, email, plus_one, is_attending, invitation_lang, invitation_id)
					VALUES($1, $2, $3, $4, $5, $6, $7, $8)`
			_, err := db.Exec(query, guest.PartyRefer, guest.FirstName, guest.LastName, guest.Email, guest.PlusOne, guest.IsAttending, guest.InvitationLang, invitationId)

			if err != nil {
				fmt.Println("error inserting guest record: ", query)
				fmt.Println(err)
				break
			}
		}

		q := `
			UPDATE guests
			SET first_name=$1, last_name=$2, email=$3, plus_one=$4, is_attending=$5, invitation_lang=$6, invitation_id = $7
			WHERE id=$8;`
		_, err := db.Exec(q, guest.FirstName, guest.LastName, guest.Email, guest.PlusOne, guest.IsAttending, guest.InvitationLang, guest.InvitationId, guest.ID)
		if err != nil {
			fmt.Println("error updating guest record: ", q)
			fmt.Println(err)
			break
		}
	}

	// delete guests that previously belongued to this party but were not part of this request
	diff := helper.FindDifference(currentGuests, requestGuests)

	if len(diff) > 0 {
		for _, i := range diff {
			query = `DELETE FROM guests WHERE id = $1`
			_, err = db.Exec(query, i)
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
