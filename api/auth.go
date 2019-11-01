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

// POST Request
func UserLogin(c echo.Context) error {
    db := db.DbManager()
    userCreds := new(model.User)
    user := new(model.User)

    // bind request to model
    if err := c.Bind(userCreds); err != nil {
        return c.JSON(http.StatusBadRequest, user)
    }

    if userCreds.Email == "" || userCreds.Password == "" {
        return c.JSON(http.StatusBadRequest, "Missing username / password")
    }

    if err := db.Where("email = ?", userCreds.Email).First(&user).Error; err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "Invalid login")
    }

    // compare password hash
    match := helper.CheckPasswordHash(userCreds.Password, user.Password)

    if match {
        tokens, err := jwttoken.GenerateTokenPair(user.ID, user.Email)
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

    db := db.DbManager()
    user := model.User{}

    signature := []byte(os.Getenv("JWT_SECRET"))
    token, err := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
        // validate the algorithm is our expected type:
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return signature, nil
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

    id := uint(claims["sub"].(float64))
    if err := db.First(&user, id).Error; err != nil {
        return err
    }

    if id == user.ID {
        newTokenPair, err := jwttoken.GenerateTokenPair(id, claims["username"].(string))
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, newTokenPair)
    } else {
        return err
    }
}
