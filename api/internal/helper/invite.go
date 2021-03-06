package helper

import (
	"fmt"
	"github.com/elkrammer/kegger/api/db"
	"github.com/elkrammer/kegger/api/model"
	"time"
)

// used by api
func FetchEventInformation(inviteId string, isProtected bool) model.Invite {
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

	// event timezone
	q = `SELECT value FROM settings WHERE name = 'time_zone';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch event time zone: %v", err)
	}
	m.TimeZone = result

	// fetch guest information
	guest := model.GuestResponse{}

	var query string
	if len(inviteId) < 10 && isProtected {
		// guestId
		query = `
			SELECT guests.*, parties.name as party_name
			FROM guests
			INNER JOIN parties ON guests.party_refer = parties.id
			WHERE guests.id = $1`
	} else if len(inviteId) > 20 && !isProtected {
		// inviteId
		query = `
			SELECT guests.*, parties.name as party_name
			FROM guests
			INNER JOIN parties ON guests.party_refer = parties.id
			WHERE guests.invitation_id = $1`
	} else {
		fmt.Printf("Invalid ID given to FetchEventInformation", err)
	}

	row = db.QueryRowx(query, inviteId)
	err = row.StructScan(&guest)
	if err != nil {
		fmt.Printf("Failed to fetch guest information from the database: %v", err)
	}
	m.Guest = guest

	// set invitation language
	m.InviteLang = guest.InvitationLang

	// event date
	q = `SELECT value FROM settings WHERE name = 'event_date';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch event date: %v", err)
	}
	layout := "2006-01-02T15:04:05Z" // iso8601
	t, _ := time.Parse(layout, result)
	loc, _ := time.LoadLocation(m.TimeZone)
	t = t.In(loc) // set location for time
	if err != nil {
		fmt.Println("Error parsing date: ", err)
	}
	// set formatting
	if m.InviteLang == "es" {
		m.EventDate = FormatSpanishDate(t)
	} else {
		m.EventDate = t.Format("Mon Jan 2 '06 at 15:04")
	}

	// get api url
	q = `SELECT value FROM settings WHERE name = 'kegger_api_url';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch kegger_api_url: %v", err)
	}
	url := result

	// invite image
	image := "invite_image_" + guest.InvitationLang
	q = `SELECT value FROM settings WHERE name = $1;`
	row = db.QueryRowx(q, image)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch invite image: %v", err)
	}
	m.InviteImage = url + result

	// signature image
	q = `SELECT value FROM settings WHERE name = 'signature_image';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch signature image: %v", err)
	}
	m.SignatureImage = url + result

	// wedding website url
	q = `SELECT value FROM settings WHERE name = 'wedding_website_url';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch website url: %v", err)
	}
	m.WeddingWebsite = result

	return m
}

// used by sendInvite function (internal)
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

	// set invitation language
	m.InviteLang = guest.InvitationLang

	// get event timezone
	q = `SELECT value FROM settings WHERE name = 'time_zone';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch event time zone: %v", err)
	}
	m.TimeZone = result

	// event date
	q = `SELECT value FROM settings WHERE name = 'event_date';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch event date: %v", err)
	}
	layout := "2006-01-02T15:04:05Z" // iso8601
	t, _ := time.Parse(layout, result)
	loc, _ := time.LoadLocation(m.TimeZone)
	t = t.In(loc) // set location for time
	if err != nil {
		fmt.Println("Error parsing date: ", err)
	}

	// set formatting
	if m.InviteLang == "es" {
		m.EventDate = FormatSpanishDate(t)
	} else {
		m.EventDate = t.Format("Mon Jan 2 '06 at 15:04")
	}

	// kegger frontend url
	q = `SELECT value FROM settings WHERE name = 'kegger_frontend_url';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch frontend: %v", err)
	}
	m.KeggerWebsite = result

	// invite image
	image := "invite_image_" + guest.InvitationLang
	q = `SELECT value FROM settings WHERE name = $1;`
	row = db.QueryRowx(q, image)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch invite image: %v", err)
	}
	m.InviteImage = result

	// wedding website url
	q = `SELECT value FROM settings WHERE name = 'wedding_website_url';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch website url: %v", err)
	}
	m.WeddingWebsite = result

	// signature image
	q = `SELECT value FROM settings WHERE name = 'signature_image';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		fmt.Printf("Failed to fetch signature image: %v", err)
	}
	m.SignatureImage = result

	return m
}
