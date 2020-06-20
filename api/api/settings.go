package api

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/elkrammer/kegger/api/db"
	"github.com/elkrammer/kegger/api/model"

	"github.com/labstack/echo/v4"
)

// GET - get all settings and values
func GetSettings(c echo.Context) error {
	db := db.DbManager()

	query := `SELECT * from settings;`
	rows, err := db.Queryx(query)

	if err != nil {
		err := fmt.Sprintf("No settings found")
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	settings := []model.Settings{}
	for rows.Next() {
		s := model.Settings{}
		err = rows.StructScan(&s)
		if err != nil {
			fmt.Println(err)
		}
		settings = append(settings, s)
	}

	return c.JSON(http.StatusOK, settings)
}

func UpdateSettings(c echo.Context) error {
	db := db.DbManager()
	settings := []model.Settings{}

	// bind request to model
	if err := c.Bind(&settings); err != nil {
		msg := fmt.Sprintf("Invalid request body. %s", err)
		return c.JSON(http.StatusBadRequest, msg)
	}

	//  update all keys as per the request
	for _, setting := range settings {
		query := `UPDATE settings SET "name" = $1, "value" = $2 WHERE "name" = $3`
		_, err := db.Exec(query, setting.Name, setting.Value, setting.Name)
		if err != nil {
			fmt.Println("error updating setting: ", query)
			fmt.Println(err)
			break
		}
	}

	return c.JSON(http.StatusOK, settings)
}

func UploadInviteImage(c echo.Context) error {
	db := db.DbManager()
	invite := c.FormValue("name")

	// source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	// destination
	filePath := "./assets/uploads/" + file.Filename
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer dst.Close()

	// copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	webPath := "/uploads/" + file.Filename
	query := `UPDATE settings SET "value" = $1 WHERE "name" = $2`
	_, err = db.Exec(query, webPath, invite)

	if err != nil {
		fmt.Println("error updating %s: ", invite, query)
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, "Invite image successfully uploaded")
}
