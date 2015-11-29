package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/samuel/go-zookeeper/zk"
)

func fetchStats(statsCache *Cache) {
	for {

		ZKservers := []string{
			// "192.168.33.10:2181",
			"192.168.99.100:32770",
			"192.168.99.100:32773",
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

}
