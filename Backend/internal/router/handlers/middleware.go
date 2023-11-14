package handlers

import (
	"github.com/breeeaaad/gproject/internal/configs"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handlers) Auth(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(401, gin.H{"msg": err})
		return
	}
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.JSON(401, gin.H{"msg": "Unexpected signing method"})
			return nil, nil
		}
		return configs.Jwtconfig(), nil
	})
	if err != nil {
		c.JSON(401, gin.H{"msg": err})
		return
	}
	if !token.Valid {
		c.JSON(401, gin.H{"msg": "Invalid token"})
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Set("id", claims["id"].(float64))
		c.Set("user", claims["user"].(string))
		c.Set("is_admin", claims["is_admin"].(bool))
	}
}
