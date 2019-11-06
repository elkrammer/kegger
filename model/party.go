package model

import (
    "time"
)

type Party struct {
        ID                  uint        `db:"id"`
        Name                string      `db:"name"`
        HostId              *uint       `db:"host_id"`
        InvitationId        *uint       `db:"invitation_id"`
        InvitationSent      *time.Time  `db:"invitation_sent"`
        InvitationOpened    *time.Time  `db:"invitation_opened"`
        IsAttending         bool        `db:"is_attending"`
        HostName            string      `db:"host_name"`
        Guests              string
        Comments            string      `db:"comments"`
}

type PartyResponse struct {
        ID                  uint        `db:"id"`
        Name                string      `db:"name"`
        HostId              *uint       `db:"host_id"`
        InvitationId        *uint       `db:"invitation_id"`
        InvitationSent      *time.Time  `db:"invitation_sent"`
        InvitationOpened    *time.Time  `db:"invitation_opened"`
        IsAttending         bool        `db:"is_attending"`
        HostName            string      `db:"host_name"`
        Guests              []Guest
        Comments            string      `db:"comments"`
}
