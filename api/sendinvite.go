package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/elkrammer/gorsvp/model"

	"github.com/labstack/echo"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// POST - send invite through sendgrid
func SendInvite(c echo.Context) error {
	email := model.SendInvite{}

	if err := c.Bind(&email); err != nil {
		msg := fmt.Sprintf("Invalid request body. %s", err)
		return c.JSON(http.StatusBadRequest, msg)
	}

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

	content := mail.NewContent("text/html", email.Message)
	m.AddContent(content)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(m)
	if err != nil {
		msg := fmt.Sprintf("There was an error sending this email to sendgrid. %s", err)
		return c.JSON(http.StatusBadRequest, msg)
	}

	return c.JSON(http.StatusOK, H{
		"msg":  "Email successfully sent",
		"code": response.StatusCode,
	})

}
