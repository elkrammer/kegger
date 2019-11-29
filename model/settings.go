package model

type Settings struct {
        Name                string          `json:"name" db:"name"`
        Value               string          `json:"value" db:"value"`
        Description         string          `json:"description" db:"description"`
}
