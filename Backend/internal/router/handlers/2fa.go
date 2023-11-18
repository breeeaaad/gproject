package handlers

import (
	"github.com/breeeaaad/gproject/internal/helpers"
	"github.com/gin-gonic/gin"
	"github.com/xlzd/gotp"
)

func (h *Handlers) Totp(c *gin.Context) {
	if user, exists := c.Get("user"); exists {
		totp := gotp.RandomSecret(16)
		if err := h.s.Authtoken(totp, user.(string)); err != nil {
			c.JSON(500, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"2fa_token": gotp.NewDefaultTOTP(totp).ProvisioningUri(user.(string), "gproject")})
	}
}

func (h *Handlers) DeleteTotp(c *gin.Context) {
	if user, exists := c.Get("user"); exists {
		var token helpers.Totp
		if err := c.BindUri(&token); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		if err := h.s.Deltoken(token.Token, user.(string)); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
	}
}
