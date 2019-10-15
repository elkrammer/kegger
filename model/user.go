package model

type User struct {
        ID                  uint        `json:"Id" gorm:"primary_key"`
        FirstName           string      `json:"firstName"`
        LastName            string      `json:"lastName"`
        Email               string      `json:"email"`
        Password            string      `json:"password"`
}
