package model

import (
    "time"
)

type Party struct {
        GuestRefer          uint
        InvitationId        uint        `json:"invitationId"`
        InvitationSent      time.Time   `json:"seenAt"`
        InvitationOpened    time.Time   `json:"seenAt"`
        IsAttending         bool        `json:"isAttending"`
        Comments            string      `json:"comments"`
}
