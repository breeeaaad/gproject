package handlers

import (
	"github.com/breeeaaad/gproject/internal/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) Login(c *gin.Context) {
	var user helpers.Account
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	id, usern, is_admin, err := h.s.Check(user)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	jwt, err := helpers.Genjwt(id, usern, is_admin)
	if err != nil {
		c.JSON(500, gin.H{"msg": "Несгенерировался токен"})
		return
	}
	c.SetCookie("jwt", jwt, 3600, "/", "localhost", true, true)
}

func (h *Handlers) Registration(c *gin.Context) {
	var user helpers.Account
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	if err := h.s.Add(user); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
}
