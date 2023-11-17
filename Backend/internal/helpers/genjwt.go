package helpers

import (
	"github.com/breeeaaad/gproject/internal/configs"
	"github.com/golang-jwt/jwt/v5"
)

func Genjwt(id int, user string, is_admin bool) (string, error) {
	prvKey, err := configs.JwtPrvKey()
	if err != nil {
		return "", err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(prvKey)
	if err != nil {
		return "", err
	}
	claims := make(jwt.MapClaims)
	claims["authorized"] = true
	claims["id"] = id
	claims["user"] = user
	claims["is_admin"] = is_admin
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}
	return token, nil
}
