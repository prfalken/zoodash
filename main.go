package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

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

	go func() {
		for {

			ZKservers := []string{
				// "192.168.33.10:2181",
				"192.168.33.10:2182",
				// "192.168.33.10:2183",
			}

			serverStats, ok := zk.FLWSrvr(ZKservers, 3*time.Second)
			statsList := []zk.ServerStats{}
			for _, s := range serverStats {
				if s.Error != nil {
					log.Debug(s.Error)
				}
				statsList = append(statsList, *s)
			}
			statsCache.lock.Lock()
			statsCache.stats = statsList
			statsCache.lock.Unlock()
			if !ok {
				log.Error("Failed to fetch stats on one or more servers")
			}

			time.Sleep(2 * time.Second)
		}

	}()

	// go func() {
	// 	for {
	// 		log.Info(statsCache.stats)

	// 		time.Sleep(2 * time.Second)
	// 	}
	// }()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		statsCache.lock.RLock()
		defer statsCache.lock.RUnlock()
		fmt.Fprintf(w, index(statsCache.stats))
	})
	http.ListenAndServe(":8080", nil)

}
