package model

type User struct {
        ID                  uint        `db:"id""`
        Name                string      `db:"name"`
        Email               string      `db:"email"`
        Password            string      `db:"password"`
}
