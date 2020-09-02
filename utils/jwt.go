package util

import (
	"time"

	"github.com/HideBa/soroha-api/config"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userId uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString([]byte(config.GetConfig().Server.KEY))
	return t
}
