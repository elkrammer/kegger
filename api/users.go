package api

import (
    "net/http"
    "fmt"

    "github.com/elkrammer/gorsvp/db"
    "github.com/elkrammer/gorsvp/model"

    "github.com/labstack/echo"
)

// GET - get user list
func GetUsers(c echo.Context) error {
    db := db.DbManager()
    users := []model.UserResponse{}

    query := `SELECT id, "name", email FROM users`
    rows, err := db.Queryx(query)

    if err != nil {
        err := fmt.Sprintf("Error fetching user list")
        return echo.NewHTTPError(http.StatusNotFound, err)
    }

    for rows.Next() {
        user := model.UserResponse{}
        err = rows.StructScan(&user)
        if err != nil {
            fmt.Println(err)
        }

        users = append(users, user)
    }

    return c.JSON(http.StatusOK, users)
}
