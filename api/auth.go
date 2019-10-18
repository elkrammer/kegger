package api

import (
    "fmt"
    "net/http"
    "os"

    "github.com/elkrammer/gorsvp/db"
    "github.com/elkrammer/gorsvp/internal/helper"
    "github.com/elkrammer/gorsvp/internal/gen_jwt_token"
    "github.com/elkrammer/gorsvp/model"

    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo"
)

func UserLogin(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    if username == "" || password == "" {
        return c.JSON(http.StatusBadRequest, "Missing username / password")
    }

    db := db.DbManager()
    user := model.User{}

    if err := db.Where("email = ?", username).First(&user).Error; err != nil {
        return err
    }

    // compare password hash
    hashed, _ := helper.HashPassword(password)
    match := helper.CheckPasswordHash(password, hashed)
    if match {
        tokens, err := jwttoken.GenerateTokenPair()
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, tokens)
    }
    return echo.ErrUnauthorized
}

func GenRefreshToken(c echo.Context) error {
    type tokenReqBody struct {
        RefreshToken string `json:"refresh_token"`
    }
    tokenReq := tokenReqBody{}
    c.Bind(&tokenReq)
    if err := c.Bind(&tokenReq); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "request body invalid")
    }

    token, err := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
        // validate the algorithm is our expected type:
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(os.Getenv("JWT_TOKEN")), nil
    })

    if err != nil {
        return err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return echo.NewHTTPError(http.StatusBadRequest, "could not renew access token")
    }
    if !token.Valid {
        return echo.NewHTTPError(http.StatusBadRequest, "refresh token invalid")
    }

    if int(claims["sub"].(float64)) == 1 {
        newTokenPair, err := jwttoken.GenerateTokenPair()
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, newTokenPair)
    } else {
        return err
    }
}
