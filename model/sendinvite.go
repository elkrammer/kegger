package model

import "time"

type SendInvite struct {
	FromName  string `json:"from_name"`
	FromEmail string `json:"from_email"`
	To        []*To  `json:"to,omitempty"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
}

type To struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Invite struct {
	EventDate     string        `json:"event_date"`
	EventLocation string        `json:"event_location"`
	EventName     string        `json:"event_name"`
	DressCode     string        `json:"dress_code"`
	Guest         GuestResponse `json:"guest"`
	Hosts         []string      `json:"hosts"`
}

type UpdateInviteRequest struct {
	InvitationId     *string    `json:"invitation_id" db:"invitation_id"`
	IsAttending      *bool      `json:"is_attending" db:"is_attending"`
	InvitationOpened *time.Time `json:"invitation_opened" db:"invitation_opened"`
}
