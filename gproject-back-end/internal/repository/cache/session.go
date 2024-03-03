package cache

import (
	"errors"
	"github.com/breeeaaad/gproject/internal/configs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	ErrNoRefreshToken = errors.New("No refresh token")
	ErrTokenExpired   = errors.New("Token expired")
)

func (c *Cache) Set(userId int, user string, isAdmin bool) (string, string, error) {
	prvKey, err := configs.JwtPrvKey()
	if err != nil {
		return "", "", err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(prvKey)
	if err != nil {
		return "", "", err
	}
	claims := make(jwt.MapClaims)
	claims["userId"] = userId
	claims["user"] = user
	claims["is_admin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	access, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", "", err
	}
	token := uuid.NewString()
	c.session[token] = refresh{
		userId:    userId,
		expiresIn: time.Now().Add(time.Hour * 60).Unix(),
		isAdmin:   isAdmin,
		username:  user,
	}
	return access, token, nil
}

func (c *Cache) Reset(uuid string) (int, string, bool, error) {
	if _, f := c.session[uuid]; !f {
		return 0, "", false, ErrNoRefreshToken
	}
	if c.session[uuid].expiresIn < time.Now().Unix() {
		delete(c.session, uuid)
		return 0, "", false, ErrTokenExpired
	}
	id, username, isAdmin := c.session[uuid].userId, c.session[uuid].username, c.session[uuid].isAdmin
	delete(c.session, uuid)
	return id, username, isAdmin, nil
}

func (c *Cache) Gout(uuid string) {
	delete(c.session, uuid)
}
