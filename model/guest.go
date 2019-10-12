package model

import (
    "time"
)

type Guest struct {
        ID          uint        `json:"id" gorm:"primary_key"`
        Name        string      `json:"name"`
        Type        string      `json:"type"`
        CreatedAt   time.Time   `json:"createdAt"`
        UpdatedAt   time.Time   `json:"updatedAt"`
}
