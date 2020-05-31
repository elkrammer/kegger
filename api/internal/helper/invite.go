package helper

import (
	"fmt"

	"github.com/elkrammer/kegger/api/db"
	"github.com/elkrammer/kegger/api/model"
)

func FetchEventInformation(inviteId string) model.Invite {
	db := db.DbManager()
	m := model.Invite{}

	// fetch event name
	var result string
	q := `SELECT value FROM settings WHERE name = 'event_name';`
	row := db.QueryRowx(q)
	err := row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch event name: %v", err)
	}
	m.EventName = result

	// event date
	q = `SELECT value FROM settings WHERE name = 'event_date';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch event date: %v", err)
	}
	m.EventDate = result

	// event location
	q = `SELECT value FROM settings WHERE name = 'event_location';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch event location: %v", err)
	}
	m.EventLocation = result

	// dress code
	q = `SELECT value FROM settings WHERE name = 'dress_code';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch dress code: %v", err)
	}
	m.DressCode = result

	// bride's name
	q = `SELECT value FROM settings WHERE name = 'bride_name';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch bride name: %v", err)
	}
	m.BrideName = result

	// groom's name
	q = `SELECT value FROM settings WHERE name = 'groom_name';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch groom name: %v", err)
	}
	m.GroomName = result

	// invite background
	q = `SELECT value FROM settings WHERE name = 'invite_background';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch invite background: %v", err)
	}
	m.InviteBackground = result

	// fetch guest information
	guest := model.GuestResponse{}

	// if the id is only a number, then it corresponds to the guest_id and not the inviteId
	query := `
		SELECT guests.*, parties.name as party_name
		FROM guests
		INNER JOIN parties ON guests.party_refer = parties.id
		WHERE guests.invitation_id = $1`
	row = db.QueryRowx(query, inviteId)
	err = row.StructScan(&guest)
	if err != nil {
		fmt.Printf("Failed to fetch guest information from the database: %v", err)
	}
	m.Guest = guest

	return m
}

func FetchEventInformationByGuestId(id int) model.Invite {
	db := db.DbManager()
	m := model.Invite{}

	// fetch event name
	var result string
	q := `SELECT value FROM settings WHERE name = 'event_name';`
	row := db.QueryRowx(q)
	err := row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch event name: %v", err)
	}
	m.EventName = result

	// event date
	q = `SELECT value FROM settings WHERE name = 'event_date';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch event date: %v", err)
	}
	m.EventDate = result

	// event location
	q = `SELECT value FROM settings WHERE name = 'event_location';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch event location: %v", err)
	}
	m.EventLocation = result

	// bride's name
	q = `SELECT value FROM settings WHERE name = 'bride_name';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch bride name: %v", err)
	}
	m.BrideName = result

	// groom's name
	q = `SELECT value FROM settings WHERE name = 'groom_name';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch groom name: %v", err)
	}
	m.GroomName = result

	// fetch guest information
	guest := model.GuestResponse{}

	// if the id is only a number, then it corresponds to the guest_id and not the inviteId
	query := `
		SELECT guests.*, parties.name as party_name
		FROM guests
		INNER JOIN parties ON guests.party_refer = parties.id
		WHERE guests.id = $1`
	row = db.QueryRowx(query, id)
	err = row.StructScan(&guest)
	if err != nil {
		fmt.Printf("Failed to fetch guest information from the database: %v", err)
	}
	m.Guest = guest

	return m
}
