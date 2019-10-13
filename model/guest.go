package model

type Guest struct {
        PartyRefer          uint
        ID                  uint        `json:"Id" gorm:"primary_key"`
        FirstName           string      `json:"firstName"`
        LastName            string      `json:"lastName"`
        Email               string      `json:"email"`
        IsAttending         bool        `json:"isAttending"`
}
