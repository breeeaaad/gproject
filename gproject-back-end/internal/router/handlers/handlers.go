package handlers

import (
	"github.com/breeeaaad/gproject/internal/repository"
	"github.com/breeeaaad/gproject/internal/repository/cache"
)

type Handlers struct {
	s *repository.Repository
	c *cache.Cache
}

func New(s *repository.Repository, c *cache.Cache) *Handlers {
	return &Handlers{
		s: s,
		c: c,
	}
}
