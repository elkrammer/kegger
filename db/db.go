package db

import (
    "fmt"
    "os"
    "github.com/elkrammer/gorsvp/internal/config"
    "github.com/elkrammer/gorsvp/model"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func Init() {
    // load config variables from .env file
    config.LoadEnv()
    conn_str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
    db, err = gorm.Open("postgres", conn_str)

    // db debugging
    db.LogMode(true)

    if err != nil {
        panic(err)
    }

    //defer db.Close()

    // migrate schema
    db.AutoMigrate(model.Guest{})
    db.AutoMigrate(model.Party{})
    db.AutoMigrate(model.User{})
    db.Model(model.Guest{}).AddForeignKey("party_refer", "parties(id)", "CASCADE", "CASCADE") // (foreign_key, destination_table, ONDELETE, ONUPDATE)

}

func DbManager() *gorm.DB {
    return db
}
