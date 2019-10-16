package main

import (
    "fmt"
    "log"
    "os"
    "strconv"

    "github.com/elkrammer/gorsvp/db"
    "github.com/elkrammer/gorsvp/internal/helper"
    "github.com/elkrammer/gorsvp/model"

    "github.com/olekukonko/tablewriter"
    "github.com/urfave/cli"
)

var app = cli.NewApp()

func createAdminUser(name, email, password string) {
    db.Init()
    db := db.DbManager()
    password, _ = helper.HashPassword(password)
    user := model.User{
        Name: name,
        Email: email,
        Password: password,
    }
    db.Create(&user)
    fmt.Sprintf("admin user %v created successfully", email)
}

func listAdminUsers() {
    db.Init()
    db := db.DbManager()
    user := model.User{}
    users := []model.User{}

    if err := db.Find(&users).Error; err != nil {
        fmt.Println(err)
    }

    if db.Find(&user).RecordNotFound() {
        fmt.Println("no admin users found")
    }

    fmt.Println("Admin User List")
    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"#", "Name", "Email"})

    for i := 1; i < len(users); i++ {
        index := strconv.Itoa(i)
        line := []string{index, users[i].Name, users[i].Email}
        table.Append(line)
    }

    table.Render()

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
                    Name:    "list",
                    Aliases: []string{"l"},
                    Usage:   "List existing administrator users",
                    Action: func(c *cli.Context) error {
                        listAdminUsers()
                        return nil
                    },
                },
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
