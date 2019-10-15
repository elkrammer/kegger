package main

import (
    "fmt"
    "os"
    "log"

    "github.com/elkrammer/gorsvp/db"
    "github.com/elkrammer/gorsvp/model"

    "github.com/urfave/cli"
)

var app = cli.NewApp()

func CreateAdminUser() {
    fmt.Println("Create admin user function!")
}

func commands() {
    app.Commands = []cli.Command{
        {
            Name:    "admin",
            Aliases: []string{"a"},
            Usage:   "Create a new Administrator User",
            Action: func(c *cli.Context) {
                CreateAdminUser()
            },
        },
    }
}

func main() {
    app.Name = "gorsvp-cmd"
    app.Usage = "Command Line Tools"
    app.Author = "Mauricio Bahamonde"
    app.Version = "1.0.0"
    commands()

    app.Action = func(c *cli.Context) error {
        fmt.Println("gorsvp-cmd - Command Line Tools")
        fmt.Println("Run gorsvp-cmd help to list available commands")
        return nil
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
