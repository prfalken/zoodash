package main

import (
	"sync"

	"github.com/samuel/go-zookeeper/zk"
)

// Cache is a simple in-memory cache for stats
type Cache struct {
	lock  sync.RWMutex
	stats map[string]zk.ServerStats
}

// NewCache contstructs a new Cache
func NewCache() *Cache {
	c := new(Cache)
	return c
}
