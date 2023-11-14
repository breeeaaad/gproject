package handlers

import "github.com/breeeaaad/gproject/internal/repository"

type Handlers struct {
	s *repository.Repository
}

func New(s *repository.Repository) *Handlers {
	return &Handlers{
		s: s,
	}
}
