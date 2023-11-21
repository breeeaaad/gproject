package handlers

import (
	"time"

	"github.com/breeeaaad/gproject/internal/helpers"
	"github.com/gin-gonic/gin"
	"github.com/xlzd/gotp"
)

func (h *Handlers) Login(c *gin.Context) {
	var user helpers.Account
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	id, usern, is_admin, token, err := h.s.Check(user)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if token != nil {
		totp := gotp.NewDefaultTOTP(token.(string))
		if !totp.Verify(user.Totp, time.Now().Unix()) {
			c.JSON(400, gin.H{"msg": "Invalid auth key"})
			return
		}
	}
	access, refresh, err := h.s.Genjwt(id, usern, is_admin)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.SetCookie("refresh", refresh, 60*60*24*7, "/main", "localhost", false, true)
	c.JSON(200, gin.H{"access": access})
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
