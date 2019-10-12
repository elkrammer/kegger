package model

type Guest struct {
        Party               []Party     `gorm:"foreignkey:GuestRefer"`
        GuestID             uint        `json:"guestId" gorm:"primary_key"`
        FirstName           string      `json:"firstName"`
        LastName            string      `json:"lastName"`
        Email               string      `json:"email"`
        IsAttending         bool        `json:"isAttending"`
}
