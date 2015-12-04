package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetLevel(log.WarnLevel)
}

func main() {

	if len(zoodashConfig.Adresses) < 1 {
		log.Error("No zookeeper servers specified in config file")
		os.Exit(1)
	}

	statsCache := NewCache()

	go fetchStats(statsCache, zoodashConfig.Adresses)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		statsCache.lock.RLock()
		defer statsCache.lock.RUnlock()
		fmt.Fprintf(w, index(statsCache.stats))
	})
	http.ListenAndServe(":8080", nil)

}
