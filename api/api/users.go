package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/elkrammer/kegger/api/db"
	"github.com/elkrammer/kegger/api/internal/helper"
	"github.com/elkrammer/kegger/api/model"

	"github.com/labstack/echo/v4"
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
	row := db.QueryRowx(query, id)
	err := row.StructScan(&user)
	if err != nil {
		err := fmt.Sprintf("User with ID: %v not found", id)
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	db := db.DbManager()
	user := model.User{}

	// bind request to model
	if err := c.Bind(&user); err != nil {
		msg := fmt.Sprintf("Invalid request body. %s", err)
		return c.JSON(http.StatusBadRequest, msg)
	}

	// request validations
	if user.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing Name field")
	}

	if user.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing Email field")
	}

	if user.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing Password field")
	}

	// check if other record with that email exists
	var exists bool
	q := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`
	err := db.QueryRow(q, user.Email).Scan(&exists)

	if exists == true {
		err := fmt.Sprintf("User with email %s already exists", user.Email)
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	hash, _ := helper.HashPassword(user.Password)
	var pid uint
	query := `
    INSERT INTO users
    (name, email, password)
    VALUES($1, $2, $3)
    RETURNING id`
	err = db.QueryRow(query, user.Name, user.Email, hash).Scan(&pid)

	if err != nil {
		fmt.Println("error inserting user record: ", query)
		fmt.Println(err)
	}

	response := model.UserResponse{}
	response.Name = user.Name
	response.Email = user.Email
	response.ID = pid
	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db := db.DbManager()

	// check if records exists
	var recordId uint
	query := `SELECT id FROM users WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&recordId)

	if err != nil && recordId == 0 {
		err := fmt.Sprintf("User with ID: %v not found", id)
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	// delete record
	query = `DELETE FROM users WHERE id = $1`
	_, err = db.Exec(query, id)

	if err != nil {
		fmt.Println("error deleting user record: ", query)
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, H{
		"deleted": id,
	})
}

func EditUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db := db.DbManager()
	user := model.User{}

	// check if records exists
	var recordId uint
	query := `SELECT id FROM users WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&recordId)

	if err != nil && recordId == 0 {
		err := fmt.Sprintf("User with ID: %v not found", id)
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	// bind request to model
	if err := c.Bind(&user); err != nil {
		msg := fmt.Sprintf("Invalid request body. %s", err)
		return c.JSON(http.StatusBadRequest, msg)
	}

	// update record that contains a new password
	if user.Password != "" {
		hash, _ := helper.HashPassword(user.Password)
		query = `
        UPDATE users
        SET name=$1, email=$2, password=$3
        WHERE id=$4;`
		_, err = db.Exec(query, user.Name, user.Email, hash, id)

		if err != nil {
			fmt.Println("error updating record: ", query)
			fmt.Println(err)
		}
		return c.JSON(http.StatusOK, user)
	}

	// record doesn't have a new password
	query = `
    UPDATE users
    SET name=$1, email=$2
    WHERE id=$3;`
	_, err = db.Exec(query, user.Name, user.Email, id)

	if err != nil {
		fmt.Println("error updating record: ", query)
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, user)
}
