package main

import (
    "github.com/elkrammer/gorsvp/db"
    "github.com/elkrammer/gorsvp/route"
)

func main() {
    db.Init()

    e := route.Init()

    // start server
    e.Logger.Fatal(e.Start(":8080"))
}
