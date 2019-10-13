package api

import (
    "net/http"
    "fmt"

    "github.com/elkrammer/gorsvp/db"
    "github.com/elkrammer/gorsvp/model"

    "github.com/labstack/echo"
)

func GetParties(c echo.Context) error {
    db := db.DbManager()

    parties := []model.Party{}
    db.Find(&parties)
    fmt.Println(db)
    fmt.Println("Endpoint Hit: getParties")

	return c.String(http.StatusOK, "Howdy!")
    //return c.JSON(http.StatusOK, parties)

    //	parties := []model.Party{}
    //	db.Find(&parties)
    // spew.Dump(json.Marshal(parties))
    // return c.JSON(http.StatusOK, parties)
    //return c.JSON(http.StatusOK, parties)
}
