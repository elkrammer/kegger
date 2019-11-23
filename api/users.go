package api

import (
    "net/http"
    "fmt"
    "strconv"

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

// GET - get individual user
func GetUser(c echo.Context) error {
    db := db.DbManager()
    user := model.UserResponse{}
    id, _ := strconv.Atoi(c.Param("id"))

    query := `SELECT id, "name", email FROM users WHERE id = $1`
    rows, err := db.Queryx(query, id)

    if err != nil {
        err := fmt.Sprintf("User with ID: %v not found", id)
        return echo.NewHTTPError(http.StatusNotFound, err)
    }

    for rows.Next() {
        err = rows.StructScan(&user)
        if err != nil {
            fmt.Println(err)
        }
    }
    return c.JSON(http.StatusOK, user)
}
