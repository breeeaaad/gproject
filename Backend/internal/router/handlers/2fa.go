package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/xlzd/gotp"
)

func (h *Handlers) TOTP(c *gin.Context) {
	if user, exists := c.Get("user"); exists {
		totp := gotp.RandomSecret(16)
		if err := h.s.Authtoken(totp, user.(string)); err != nil {
			c.JSON(500, gin.H{"msg": err})
		}
		c.JSON(200, gin.H{"2fa_token": gotp.NewDefaultTOTP(totp).ProvisioningUri(user.(string), "gproject")})
	}
}
