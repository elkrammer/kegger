package model

import (
	"gopkg.in/guregu/null.v3"
)

type Guest struct {
	ID               int         `json:"id" db:"id"`
	PartyRefer       int         `json:"party_refer" db:"party_refer"`
	FirstName        string      `json:"first_name" db:"first_name"`
	LastName         string      `json:"last_name" db:"last_name"`
	InvitationId     null.String `json:"invitation_id" db:"invitation_id"`
	InvitationLang   string      `json:"invitation_lang" db:"invitation_lang"`
	InvitationSent   null.Time   `json:"invitation_sent" db:"invitation_sent"`
	InvitationOpened null.Time   `json:"invitation_opened" db:"invitation_opened"`
	Email            string      `json:"email" db:"email"`
	PlusOne          bool        `json:"plus_one" db:"plus_one"`
	IsAttending      bool        `json:"is_attending" db:"is_attending"`
}

type GuestResponse struct {
	PartyName        string    `json:"party_name" db:"party_name"`
	PartyRefer       int       `json:"party_refer" db:"party_refer"`
	ID               int       `json:"id" db:"id"`
	FirstName        string    `json:"first_name" db:"first_name"`
	LastName         string    `json:"last_name" db:"last_name"`
	InvitationId     string    `json:"invitation_id" db:"invitation_id"`
	InvitationLang   string    `json:"invitation_lang" db:"invitation_lang"`
	InvitationSent   null.Time `json:"invitation_sent" db:"invitation_sent"`
	InvitationOpened null.Time `json:"invitation_opened" db:"invitation_opened"`
	Email            string    `json:"email" db:"email"`
	PlusOne          bool      `json:"plus_one" db:"plus_one"`
	IsAttending      bool      `json:"is_attending" db:"is_attending"`
}
