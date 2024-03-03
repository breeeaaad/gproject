package cache

import (
	"sync"
)

type refresh struct {
	userId    int
	expiresIn int64
	isAdmin   bool
	username  string
}
type Cache struct {
	sync.RWMutex
	session map[string]refresh
}

func New() *Cache {
	return &Cache{
		session: make(map[string]refresh, 16),
	}
}
