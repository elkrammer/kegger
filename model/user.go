package model

type User struct {
    ID                  uint        `db:"id"`
    Name                string      `db:"name" json:"name"`
    Email               string      `db:"email" json:"email"`
    Password            string      `db:"password" json:"password"`
}

type UserResponse struct {
    ID                  uint        `db:"id" json:"id"`
    Name                string      `db:"name" json:"name"`
    Email               string      `db:"email" json:"email"`
}
