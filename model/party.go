package model

import (
    "time"
    "gopkg.in/guregu/null.v3"
)

type Party struct {
        ID                  uint           `db:"id"`
        Name                string         `db:"name"`
        HostId              null.Int       `db:"host_id"`
        InvitationId        null.Int       `db:"invitation_id"`
        InvitationSent      null.Time      `db:"invitation_sent"`
        InvitationOpened    null.Time      `db:"invitation_opened"`
        IsAttending         null.Bool      `db:"is_attending"`
        HostName            null.String    `db:"host_name"`
        Guests              string
        Comments            null.String    `db:"comments"`
}

type PartyResponse struct {
        ID                  uint           `db:"id"`
        Name                string         `db:"name"`
        InvitationId        null.Int       `db:"invitation_id"`
        InvitationSent      null.Time      `db:"invitation_sent"`
        InvitationOpened    null.Time      `db:"invitation_opened"`
        IsAttending         null.Bool      `db:"is_attending"`
        HostName            null.String    `db:"host_name"`
        HostId              null.Int       `db:"host_id"`
        Guests              []Guest
        Comments            null.String    `db:"comments"`
}

type PartyRequest struct {
        ID                  uint        `db:"id"`
        Name                string      `db:"name"`
        InvitationId        *uint       `db:"invitation_id"`
        InvitationSent      *time.Time  `db:"invitation_sent"`
        InvitationOpened    *time.Time  `db:"invitation_opened"`
        IsAttending         bool        `db:"is_attending"`
        HostId              *uint       `db:"host_id"`
        Guests              []Guest
        Comments            string      `db:"comments"`
}
