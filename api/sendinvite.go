package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/elkrammer/gorsvp/db"
	"github.com/elkrammer/gorsvp/internal/helper"
	"github.com/elkrammer/gorsvp/model"

	"github.com/labstack/echo/v4"
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

	// fetch the data we need to populate the invites
	invite := helper.FetchEventInformationByGuestId(id)

	body := helper.ProcessTemplateFile("templates/invite.tpl", invite)

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

	// update invite_sent column
	query := `UPDATE guests SET invitation_sent = NOW() WHERE id=$1;`
	_, err = db.Exec(query, id)
	if err != nil {
		msg := fmt.Sprintf("There was an error updating invitation_sent column for guest %v: %v", id, err)
		return c.JSON(http.StatusBadRequest, msg)
	}

	return c.JSON(http.StatusOK, H{
		"msg":  "Email successfully sent",
		"code": response.StatusCode,
	})
}

func FetchInviteDetails(c echo.Context) error {
	id := c.Param("invite_id")

	if len(id) == 0 {
		msg := fmt.Sprintf("Invite ID can't be null")
		return c.JSON(http.StatusBadRequest, msg)
	}

	invite := helper.FetchEventInformation(id)

	// validate if we were able to fetch invite details
	if invite.Guest.PartyName == "" {
		msg := fmt.Sprintf("Could not fetch invite for id: %v", id)
		return c.JSON(http.StatusBadRequest, msg)
	}
	return c.JSON(http.StatusOK, invite)
}
