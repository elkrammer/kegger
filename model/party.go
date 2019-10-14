package model

import (
    "time"
)

type Party struct {
        ID                  uint        `json:"Id" gorm:"primary_key"`
        Name                string      `json:"name"`
        Guests              []Guest     `gorm:"foreignkey:PartyRefer"`
        InvitationId        *uint       `json:"invitationId"`
        InvitationSent      *time.Time  `json:"invitationSent"`
        InvitationOpened    *time.Time  `json:"invitationOpened"`
        IsAttending         bool        `json:"isAttending"`
        Comments            string      `json:"comments"`
}
