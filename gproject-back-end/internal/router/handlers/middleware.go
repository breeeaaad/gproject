package handlers

import (
	"errors"

	"github.com/breeeaaad/gproject/internal/configs"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handlers) Auth(c *gin.Context) {
	cookie, err := c.Cookie("access")
	if err != nil {
		if err := h.Resetoken(c); err != nil {
			c.JSON(401, gin.H{"msg": err.Error()})
			return
		}
		h.Auth(c)
		return
	}
	pubKey, err := configs.JwtPubKey()
	if err != nil {
		c.JSON(401, gin.H{"msg": err.Error()})
		return
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		c.JSON(401, gin.H{"msg": err.Error()})
		return
	}
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			c.JSON(401, gin.H{"msg": "Unexpected signing method"})
			return nil, nil
		}
		return key, nil
	})
	if err != nil {
		if err := h.Resetoken(c); err != nil {
			c.JSON(401, gin.H{"msg": err.Error()})
			return
		}
		h.Auth(c)
		return
	}
	if !token.Valid {
		c.JSON(401, gin.H{"msg": "Invalid access token"})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		c.Set("id", claims["id"].(float64))
		c.Set("user", claims["user"].(string))
		c.Set("is_admin", claims["is_admin"].(bool))
	}
}

func (h *Handlers) Resetoken(c *gin.Context) error {
	cookie, err := c.Cookie("refresh")
	if err != nil {
		return err
	}
	id, user, is_admin, err := h.s.Refresh(cookie)
	if err != nil {
		return err
	}
	check, err := h.s.DelRefresh(cookie)
	if err != nil {
		return err
	}
	if !check {
		return errors.New("Token expired")
	}
	access, refresh, err := h.s.Genjwt(id, user, is_admin)
	if err != nil {
		return err
	}
	c.SetCookie("refresh", refresh, 60*60*24*7, "/main", "localhost", false, true)
	c.SetCookie("access", access, 60*15, "/main", "localhost", false, true)
	return nil
}
