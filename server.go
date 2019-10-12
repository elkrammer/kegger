package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/elkrammer/gorsvp/model"
    "github.com/elkrammer/gorsvp/internal/config"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    // load config variables from .env file
    config.LoadEnv()

    // echo instance
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // create db connection
    conn_str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
                    os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"),
                    os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
    db, err := gorm.Open("postgres", conn_str)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // migrate schema
    db.AutoMigrate(model.Guest{})
    db.AutoMigrate(model.Party{})

    // routes
    e.GET("/", hello)

    // start server
    e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
}
