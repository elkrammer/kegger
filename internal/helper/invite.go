package helper

import (
	"fmt"

	"github.com/elkrammer/gorsvp/db"
	"github.com/elkrammer/gorsvp/model"
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

	// hosts
	var hosts []string
	q = `SELECT DISTINCT users.name AS host_name FROM parties INNER JOIN users ON parties.host_id = users.id;`
	rows, err := db.Queryx(q)

	if err != nil {
		fmt.Printf("Failed to fetch hosts: %v", err)
	}

	for rows.Next() {
		var host string
		if err := rows.Scan(&host); err != nil {
			fmt.Printf("%v", err)
		}
		hosts = append(hosts, host)
	}
	m.Hosts = hosts

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

	// hosts
	var hosts []string
	q = `SELECT DISTINCT users.name AS host_name FROM parties INNER JOIN users ON parties.host_id = users.id;`
	rows, err := db.Queryx(q)

	if err != nil {
		fmt.Printf("Failed to fetch hosts: %v", err)
	}

	for rows.Next() {
		var host string
		if err := rows.Scan(&host); err != nil {
			fmt.Printf("%v", err)
		}
		hosts = append(hosts, host)
	}
	m.Hosts = hosts

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
