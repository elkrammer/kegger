package model

import (
    "gopkg.in/guregu/null.v3"
)

type Party struct {
        ID                  uint           `db:"id"`
        Name                string         `db:"name"`
        HostId              null.Int       `db:"host_id"`
        HostName            null.String    `db:"host_name"`
        Guests              string
        Comments            null.String    `db:"comments"`
}

type PartyResponse struct {
        ID                  uint           `json:"id" db:"id"`
        Name                string         `json:"name" db:"name"`
        HostId              null.Int       `json:"host_id" db:"host_id"`
        HostName            null.String    `json:"host_name" db:"host_name"`
        Guests              []Guest
        Comments            null.String    `json:"comments" db:"comments"`
}

type PartyRequest struct {
        ID                  uint        `json:"id" db:"id"`
        Name                string      `json:"name" db:"name"`
        HostId              *uint       `json:"host_id" db:"host_id"`
        IsAttending         bool        `json:"is_attending" db:"is_attending"`
        Guests              []Guest
        Comments            string      `json:"comments" db:"comments"`
}
