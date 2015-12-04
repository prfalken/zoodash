package main

import (
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/samuel/go-zookeeper/zk"
)

func fetchStats(statsCache *Cache, zkServers map[string]string) {

	for {
		hosts := []string{}
		ipPorts := []string{}

		for _, ipPort := range zkServers {
			ipPorts = append(ipPorts, ipPort)
		}

		for host := range zkServers {
			hosts = append(hosts, host)
		}

		serverStats, ok := zk.FLWSrvr(ipPorts, 3*time.Second)

		stats := map[string]zk.ServerStats{}
		for idx, s := range serverStats {
			if s.Error != nil {
				log.Debug(s.Error)
			}
			stats[hosts[idx]] = *s
		}
		statsCache.lock.Lock()
		statsCache.stats = stats
		statsCache.lock.Unlock()
		if !ok {
			log.Error("Failed to fetch stats on one or more servers")
		}

		time.Sleep(2 * time.Second)
	}

}
