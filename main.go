package main

import (
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {

	if len(zoodashConfig.Adresses) < 1 {
		log.Error("No zookeeper servers specified in config file")
		os.Exit(1)
	}

	statsCache := NewStats()

	go statsCache.fetchStats(zoodashConfig.Adresses)

	router := httprouter.New()
	router.GET("/", indexHandler)
	router.GET("/api/", apiHandler)
	router.GET("/browse", browseHandler)
	router.GET("/stats", statsCache.statsHandler)

	log.Info("Listening...")
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, router)))

}
