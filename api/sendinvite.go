package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/elkrammer/gorsvp/db"
	"github.com/elkrammer/gorsvp/internal/helper"
	"github.com/elkrammer/gorsvp/model"

	"github.com/labstack/echo"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// POST - send invite through sendgrid
func SendInvite(c echo.Context) error {
	db := db.DbManager()
	email := model.SendInvite{}
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&email); err != nil {
		msg := fmt.Sprintf("Invalid request body. %s", err)
		return c.JSON(http.StatusBadRequest, msg)
	}

	// fetch all the data we need to prepopulate the invites
	vars := make(map[string]interface{})

	// event name
	var result string
	q := `SELECT value FROM settings WHERE name = 'event_name';`
	row := db.QueryRowx(q)
	err := row.Scan(&result)
	if err != nil {
		e := fmt.Sprintf("Failed to fetch event information from the database: %v", err)
		return echo.NewHTTPError(http.StatusNotFound, e)
	}
	vars["EventName"] = result

	// event date
	q = `SELECT value FROM settings WHERE name = 'event_date';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		e := fmt.Sprintf("Failed to fetch event information from the database: %v", err)
		return echo.NewHTTPError(http.StatusNotFound, e)
	}
	vars["EventDate"] = result

	// event location
	q = `SELECT value FROM settings WHERE name = 'event_location';`
	row = db.QueryRowx(q)
	err = row.Scan(&result)
	if err != nil {
		e := fmt.Sprintf("Failed to fetch event information from the database: %v", err)
		return echo.NewHTTPError(http.StatusNotFound, e)
	}
	vars["EventLocation"] = result

	// hosts
	var hosts []string
	q = `SELECT DISTINCT users.name AS host_name FROM parties INNER JOIN users ON parties.host_id = users.id;`
	rows, err := db.Queryx(q)

	if err != nil {
		err := fmt.Sprintf("Error fetching host list")
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	for rows.Next() {
		var host string
		if err := rows.Scan(&host); err != nil {
			fmt.Printf("%v", err)
		}
		hosts = append(hosts, host)
	}
	vars["Hosts"] = hosts

	// fetch guest information
	guest := model.GuestResponse{}
	query := `
	SELECT guests.*, parties.name as party_name
	FROM guests
	INNER JOIN parties ON guests.party_refer = parties.id
	where guests.id = $1`
	row = db.QueryRowx(query, id)
	err = row.StructScan(&guest)
	if err != nil {
		e := fmt.Sprintf("Failed to fetch guest information from the database: %v", err)
		return echo.NewHTTPError(http.StatusNotFound, e)
	}
	vars["Name"] = guest.FirstName
	vars["InviteId"] = guest.InvitationId

	body := helper.ProcessTemplateFile("templates/invite.tpl", vars)

	m := mail.NewV3Mail()
	from := mail.NewEmail(email.FromName, email.FromEmail)
	m.SetFrom(from)
	m.Subject = email.Subject

	p := mail.NewPersonalization()

	var recipientList []*mail.Email
	for _, r := range email.To {
		recipientList = append(recipientList, mail.NewEmail(r.Name, r.Email))
	}

	p.AddTos(recipientList...)
	m.AddPersonalizations(p)

	co := mail.NewContent("text/html", body)
	m.AddContent(co)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(m)
	if err != nil {
		msg := fmt.Sprintf("There was an error sending this email to sendgrid. %s", err)
		return c.JSON(http.StatusBadRequest, msg)
	}

	if response.StatusCode >= 400 {
		msg := fmt.Sprintf("There was an error sending the email to sendgrid. SendGrid returned Status Code: %v. Status Message: %v", response.StatusCode, response.Body)
		return c.JSON(http.StatusBadRequest, msg)
	}

	// TODO: set invite as sent in db

	return c.JSON(http.StatusOK, H{
		"msg":  "Email successfully sent",
		"code": response.StatusCode,
	})

}
