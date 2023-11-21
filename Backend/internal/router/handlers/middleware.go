package handlers

import (
	"time"

	"github.com/breeeaaad/gproject/internal/configs"
	"github.com/breeeaaad/gproject/internal/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handlers) Auth(c *gin.Context) {
	var access helpers.Request
	if err := c.BindJSON(&access); err != nil {
		c.JSON(401, gin.H{"msg": err.Error()})
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
	token, err := jwt.Parse(access.Access, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			c.JSON(401, gin.H{"msg": "Unexpected signing method"})
			return nil, nil
		}
		return key, nil
	})
	if err != nil {
		c.JSON(401, gin.H{"msg": err.Error()})
		return
	}
	if !token.Valid {
		c.JSON(401, gin.H{"msg": "Invalid access token"})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		if exp, boo := claims["exp"].(int64); boo {
			if exp < time.Now().Unix() {
				c.Set("id", claims["id"].(float64))
				c.Set("user", claims["user"].(string))
				c.Set("is_admin", claims["is_admin"].(bool))
				return
			}
		}
	}
	cookie, err := c.Cookie("refresh")
	if err != nil {
		c.JSON(401, gin.H{"msg": err.Error()})
		return
	}
	check, err := h.s.Refresh(cookie)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if !check {
		c.JSON(401, gin.H{"msg": "Token expired"})
		return
	}
	if id, ok := claims["id"].(int); ok {
		access, refresh, err := h.s.Genjwt(id, claims["user"].(string), claims["is_admin"].(bool))
		if err != nil {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		c.SetCookie("refresh", refresh, 60*60*24*7, "/main", "localhost", false, true)
		c.JSON(200, gin.H{"access": access})
	}
}
