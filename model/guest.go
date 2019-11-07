package model

type Guest struct {
        PartyRefer          int         `json:"party_refer"`
        ID                  int         `json:"id"`
        FirstName           string      `json:"first_name"`
        LastName            string      `json:"last_name"`
        Email               string      `json:"email"`
        IsAttending         bool        `json:"is_attending"`
}

type GuestResponse struct {
        PartyName           string      `json:"party_name" db:"party_name"`
        PartyRefer          int         `json:"-" db:"party_refer"`
        ID                  int         `json:"id"`
        FirstName           string      `json:"first_name" db:"first_name"`
        LastName            string      `json:"last_name" db:"last_name"`
        Email               string      `json:"email" db:"email"`
        IsAttending         bool        `json:"is_attending" db:"is_attending"`
}
