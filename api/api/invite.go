package api

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/elkrammer/kegger/api/db"
	"github.com/elkrammer/kegger/api/internal/helper"
	"github.com/elkrammer/kegger/api/model"

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

	// process template based on guest's preferred language
	template_file := fmt.Sprintf("templates/invite_%s.tpl", invite.InviteLang)
	body := helper.ProcessTemplateFile(template_file, invite)

	m := mail.NewV3Mail()
	from := mail.NewEmail(email.FromName, email.FromEmail)
	m.SetFrom(from)

	// set subject
	if invite.InviteLang == "es" {
		m.Subject = fmt.Sprintf("Te invitamos a celebrar el matrimonio de %s", invite.EventName)
	}

	if invite.InviteLang == "en" {
		m.Subject = fmt.Sprintf("Join us to celebrate the wedding of %s", invite.EventName)
	}

	p := mail.NewPersonalization()

	var recipientList []*mail.Email
	for _, r := range email.To {
		recipientList = append(recipientList, mail.NewEmail(r.Name, r.Email))
	}

	p.AddTos(recipientList...)
	m.AddPersonalizations(p)

	// read/attach invite image
	image := "./assets" + invite.InviteImage
	a_jpg := mail.NewAttachment()
	dat, err := ioutil.ReadFile(image)
	if err != nil {
		msg := fmt.Sprintf("Unable to read invite image %s", err)
		fmt.Println(msg)
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(dat))
	a_jpg.SetContent(encoded)
	a_jpg.SetType("image/jpeg")
	a_jpg.SetFilename("invite.jpg")
	a_jpg.SetDisposition("attachment")
	a_jpg.SetContentID("Invite")

	co := mail.NewContent("text/html", body)
	m.AddContent(co)
	m.AddAttachment(a_jpg)

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

	invite := helper.FetchEventInformation(id, false)

	// validate if we were able to fetch invite details
	if invite.Guest.PartyName == "" {
		msg := fmt.Sprintf("Could not fetch invite for id: %v", id)
		return c.JSON(http.StatusBadRequest, msg)
	}
	return c.JSON(http.StatusOK, invite)
}

func FetchInviteDetailsProtected(c echo.Context) error {
	id := c.Param("guest_id")

	if len(id) == 0 {
		msg := fmt.Sprintf("Invite ID can't be null")
		return c.JSON(http.StatusBadRequest, msg)
	}

	invite := helper.FetchEventInformation(id, true)

	// validate if we were able to fetch invite details
	if invite.Guest.PartyName == "" {
		msg := fmt.Sprintf("Could not fetch invite for id: %v", id)
		return c.JSON(http.StatusBadRequest, msg)
	}
	return c.JSON(http.StatusOK, invite)
}

func UpdateInvite(c echo.Context) error {
	db := db.DbManager()
	request := model.UpdateInviteRequest{}

	if err := c.Bind(&request); err != nil {
		msg := fmt.Sprintf("Invalid request body. %s", err)
		return c.JSON(http.StatusBadRequest, msg)
	}

	guest := model.Guest{}
	query := `
	SELECT *
	FROM guests
	where invitation_id = $1`
	err := db.QueryRowx(query, request.InvitationId).StructScan(&guest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Invitation ID")
	}

	// update is_attending field
	query = `UPDATE guests SET is_attending = $1, plus_one = $2 WHERE invitation_id = $3`
	_, err = db.Exec(query, request.IsAttending, request.PlusOne, request.InvitationId)
	if err != nil {
		msg := fmt.Sprintf("There was an error updating invitation_sent column for guest %v: %v", guest.ID, err)
		return c.JSON(http.StatusBadRequest, msg)
	}

	// update invitation_opened field
	query = `UPDATE guests SET invitation_opened = $1 WHERE invitation_id = $2`
	_, err = db.Exec(query, request.InvitationOpened, request.InvitationId)
	if err != nil {
		msg := fmt.Sprintf("There was an error updating invitation_opened column for guest %v: %v", guest.ID, err)
		return c.JSON(http.StatusBadRequest, msg)
	}

	msg := fmt.Sprintf("Successfully updated invite for guest %s", guest.FirstName)
	return c.JSON(http.StatusOK, msg)
}

// GET - find invitation id from email
func FindInviteId(c echo.Context) error {
	email := c.Param("email")
	db := db.DbManager()

	if len(email) == 0 {
		msg := fmt.Sprintf("Email can't be null")
		return c.JSON(http.StatusBadRequest, msg)
	}

	// attempt to fetch invite id
	var inviteId string
	q := `SELECT invitation_id FROM guests WHERE email = $1`
	row := db.QueryRowx(q, email)
	err := row.Scan(&inviteId)
	if err != nil {
		msg := fmt.Sprintf("No invitations found for email: %v", email)
		return c.JSON(http.StatusBadRequest, msg)
	}

	return c.JSON(http.StatusOK, inviteId)
}
