package helpers

import (
	"github.com/breeeaaad/gproject/internal/configs"
	"github.com/golang-jwt/jwt/v5"
)

func Genjwt(id int, user string, is_admin bool) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["id"] = id
	claims["user"] = user
	claims["is_admin"] = is_admin
	jwtoken, err := token.SignedString(configs.Jwtconfig())
	return jwtoken, err
}
