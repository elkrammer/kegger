package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/elkrammer/kegger/api/db"
	jwttoken "github.com/elkrammer/kegger/api/internal/gen_jwt_token"
	"github.com/elkrammer/kegger/api/internal/helper"
	"github.com/elkrammer/kegger/api/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// POST Request
func UserLogin(c echo.Context) error {
	db := db.DbManager()
	request := new(model.User)

	// bind request to model
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	User := new(model.User)
	query := "SELECT id, email, password FROM users WHERE email = $1"
	err := db.QueryRowx(query, request.Email).StructScan(User)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Invalid login")
	}

	// compare password hash
	match := helper.CheckPasswordHash(request.Password, User.Password)

	if match {
		tokens, err := jwttoken.GenerateTokenPair(User.ID, User.Email)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, tokens)
	}
	return echo.ErrUnauthorized
}

func GenRefreshToken(c echo.Context) error {
	type tokenReqBody struct {
		AccessToken string `json:"access_token"`
	}

	tokenReq := tokenReqBody{}
	c.Bind(&tokenReq)

	db := db.DbManager()

	signature := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenReq.AccessToken, func(token *jwt.Token) (interface{}, error) {
		// validate the algorithm is our expected type:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signature, nil
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error parsing jwt token - could not renew access token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "could not renew access token")
	}
	if !token.Valid {
		return echo.NewHTTPError(http.StatusBadRequest, "refresh token invalid")
	}

	var UserId uint
	id := uint(claims["sub"].(float64))
	row := db.QueryRow("SELECT id FROM users WHERE users.id = $1", id)
	err = row.Scan(&UserId)

	if err != nil {
		return err
	}

	if id == UserId {
		newTokenPair, err := jwttoken.GenerateTokenPair(id, claims["username"].(string))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, newTokenPair)
	}

	return err
}
