package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	conn    *pgx.Conn
	context context.Context
}

func New(conn *pgx.Conn) *Repository {
	return &Repository{
		conn:    conn,
		context: context.Background(),
	}
}
