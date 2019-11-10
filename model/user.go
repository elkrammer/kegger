package model

type User struct {
    ID                  uint        `db:"id"`
    Name                string      `db:"name"`
    Email               string      `db:"email"`
    Password            string      `db:"password"`
}

type UserResponse struct {
    ID                  uint        `db:"id" json:"id"`
    Name                string      `db:"name" json:"name"`
    Email               string      `db:"email" json:"email"`
}
