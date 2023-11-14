package main

import (
	"context"

	"github.com/breeeaaad/gproject/internal/configs"
	"github.com/breeeaaad/gproject/internal/repository"
	"github.com/breeeaaad/gproject/internal/router"
	"github.com/breeeaaad/gproject/internal/router/handlers"
)

func main() {
	c := configs.Dbconfig()
	defer c.Close(context.Background())
	r := repository.New(c)
	h := handlers.New(r)
	router.Router(h)
}
