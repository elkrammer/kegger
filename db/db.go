package db

import (
    "fmt"
    "os"
    _ "github.com/jackc/pgx/stdlib"
    "github.com/jmoiron/sqlx"
    "github.com/elkrammer/gorsvp/internal/config"
    //    "github.com/elkrammer/gorsvp/model"
)

var db *sqlx.DB
var err error

func Init() {
    // load config variables from .env file
    config.LoadEnv()
    connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

    db = sqlx.MustConnect("pgx", connectionString)
    err = db.Ping()
    if err != nil {
        fmt.Println("Failed to connect to database: " + connectionString)
    }
}

func DbManager() *sqlx.DB {
    if db == nil {
        panic("Failed to get DB Connection")
    }
    return db
}
