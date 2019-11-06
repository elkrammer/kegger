package model

type Guest struct {
        PartyRefer          int         `json:"party_refer"`
        ID                  int         `json:"id"`
        FirstName           string      `json:"first_name"`
        LastName            string      `json:"last_name"`
        Email               string      `json:"email"`
        IsAttending         bool        `json:"is_attending"`
}
