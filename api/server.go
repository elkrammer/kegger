package main

import (
	"github.com/elkrammer/kegger/api/db"
	"github.com/elkrammer/kegger/api/route"
)

func main() {
	db.Init()

	e := route.Init()

	// start server
	e.Logger.Fatal(e.Start(":4040"))
}
