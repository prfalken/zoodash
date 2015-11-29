package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/samuel/go-zookeeper/zk"
)

func fetchStats(statsCache *Cache, zkServers []string) {
	for {

		serverStats, ok := zk.FLWSrvr(zkServers, 3*time.Second)
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

}
