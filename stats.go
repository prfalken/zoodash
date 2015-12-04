package main

import (
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/samuel/go-zookeeper/zk"
)

type StatsStorage struct {
	lock  sync.RWMutex
	stats map[string]zk.ServerStats
}

// NewStats contstructs a new StatsStorage
func NewStats() *StatsStorage {
	c := new(StatsStorage)
	return c
}

func (store *StatsStorage) fetchStats(zkServers map[string]string) {

	for {
		hosts := getZKHostsFromConfig()
		ipPorts := []string{}
		// keep order when splitting in hosts list and ip:port list
		for _, h := range hosts {
			ipPorts = append(ipPorts, zkServers[h])
		}
		servers, ok := zk.FLWSrvr(ipPorts, 3*time.Second)

		recorded := map[string]zk.ServerStats{}
		for idx, s := range servers {
			if s.Error != nil {
				log.Debug(s.Error)
			}
			recorded[hosts[idx]] = *s
		}
		store.lock.Lock()
		store.stats = recorded
		store.lock.Unlock()
		if !ok {
			log.Error("Failed to fetch stats on one or more servers")
		}

		time.Sleep(2 * time.Second)
	}

}

func (s *StatsStorage) getStats() map[string]zk.ServerStats {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.stats
}
