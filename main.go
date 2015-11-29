package main

import (
	"fmt"
	"net/http"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/samuel/go-zookeeper/zk"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

// Cache is a simple in-memory cache for stats
type Cache struct {
	lock  sync.RWMutex
	stats []zk.ServerStats
}

// NewCache contstructs a new Cache
func NewCache() *Cache {
	c := new(Cache)
	return c
}

func main() {
	statsCache := NewCache()

	go fetchStats(statsCache)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		statsCache.lock.RLock()
		defer statsCache.lock.RUnlock()
		fmt.Fprintf(w, index(statsCache.stats))
	})
	http.ListenAndServe(":8080", nil)

}
