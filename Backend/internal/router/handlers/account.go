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
		c.JSON(400, gin.H{"msg": err})
		return
	}
	id, usern, is_admin, token, err := h.s.Check(user)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	if token != nil {
		totp := gotp.NewDefaultTOTP(token.(string))
		if !totp.Verify(user.Totp, time.Now().Unix()) {
			c.JSON(400, gin.H{"msg": "Invalid auth key"})
			return
		}
	}
	jwt, err := helpers.Genjwt(id, usern, is_admin)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}
	c.SetCookie("jwt", jwt, 3600, "/", "localhost", false, true)
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
