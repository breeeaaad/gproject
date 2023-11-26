package router

import (
	"github.com/breeeaaad/gproject/internal/router/handlers"
	"github.com/gin-gonic/gin"
)

func Router(h *handlers.Handlers) {
	r := gin.Default()
	auth := r.Group("/")
	{
		auth.POST("/registration", h.Registration)
		auth.POST("/login", h.Login)
		authorized := auth.Group("/main", h.Auth)
		{
			authorized.PUT("/2fa", h.Totp)
			authorized.DELETE("/:token", h.DeleteTotp)
			authorized.DELETE("/logout", h.Logout)
		}
	}
	r.Run()
}
