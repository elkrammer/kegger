package model

import (
    "time"
)

type Party struct {
        ID                  uint        `json:"Id" gorm:"primary_key"`
        Guests              []Guest     `gorm:"foreignkey:PartyRefer"`
        InvitationId        uint        `json:"invitationId"`
        InvitationSent      time.Time   `json:"seenAt"`
        InvitationOpened    time.Time   `json:"seenAt"`
        IsAttending         bool        `json:"isAttending"`
        Comments            string      `json:"comments"`
}
