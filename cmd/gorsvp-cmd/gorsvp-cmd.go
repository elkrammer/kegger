package main

import (
    "fmt"
    "os"
    "log"
    "golang.org/x/crypto/bcrypt"

    "github.com/elkrammer/gorsvp/db"
    "github.com/elkrammer/gorsvp/model"

    "github.com/urfave/cli"
)

var app = cli.NewApp()

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func createAdminUser(name, email, password string) {
    db.Init()
    db := db.DbManager()
    password, _ = HashPassword(password)
    user := model.User{
        Name: name,
        Email: email,
        Password: password,
    }
    db.Create(&user)
    fmt.Sprintf("admin user %v created successfully", email)
}

func deleteAdminUser() {
    fmt.Println("delete admin user")
}

func commands() {
    app.Commands = []cli.Command{
        {
            Name:    "admin",
            Aliases: []string{"a"},
            Usage:   "gorsvp-cmd admin --help" ,
            Subcommands: []cli.Command{
                {
                    Name:    "add",
                    Aliases: []string{"a"},
                    Usage:   "Create a new administrator user",
                    Action: func(c *cli.Context) error {
                        name := c.String("n")
                        email := c.String("e")
                        password := c.String("p")
                        if name != "" && email != "" && password != "" {
                            createAdminUser(name, email, password)
                        } else {
                            fmt.Println("missing n, e and p flags for the add operation")
                        }
                        return nil
                    },
                    Flags: []cli.Flag{
                        cli.StringFlag{
                            Name: "name, n",
                        },
                        cli.StringFlag{
                            Name: "email, e",
                        },
                        cli.StringFlag{
                            Name: "password, p",
                        },
                    },
                },
                {
                    Name:    "delete",
                    Aliases: []string{"d"},
                    Usage:   "Delete an existing administrator user",
                    Action: func(c *cli.Context) error {
                        deleteAdminUser()
                        return nil
                    },
                },
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
