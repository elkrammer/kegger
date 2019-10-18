package model

type User struct {
        ID                  uint        `json:"Id" gorm:"primary_key"`
        Name                string      `json:"Name"`
        Email               string      `json:"email"`
        Password            string      `json:"password"`
}
