package jwttoken

import (
  "time"
  "os"
  "github.com/dgrijalva/jwt-go"
)

func GenerateTokenPair(username string) (map[string]string, error) {
  token := jwt.New(jwt.SigningMethodHS256)
  claims := token.Claims.(jwt.MapClaims)
  claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
  claims["sub"] = 1
  claims["username"] = username

  secret := os.Getenv("JWT_SECRET")

  // generate encoded token and send it as response.
  t, err := token.SignedString([]byte(secret))
  if err != nil {
    return nil, err
  }

  refreshToken := jwt.New(jwt.SigningMethodHS256)
  rtClaims := refreshToken.Claims.(jwt.MapClaims)
  rtClaims["sub"] = 1
  rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

  rt, err := refreshToken.SignedString([]byte(secret))
  if err != nil {
    return nil, err
  }

  return map[string]string{
    "access_token":  t,
    "refresh_token": rt,
  }, nil

}
