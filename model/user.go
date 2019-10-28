package model

type User struct {
        ID                  uint        `json:"Id" gorm:"primary_key"`
        Name                string      `json:"Name"`
        Email               string      `json:"email" gorm:"unique_index"`
        Password            string      `json:"password"`
}
