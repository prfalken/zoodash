package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/samuel/go-zookeeper/zk"
)

func ZKBrowse(path string) []string {
	ipPorts := getZKAddressesFromConfig()
	conn, _, err := zk.Connect(ipPorts, 2*time.Second)
	if err != nil {
		log.Error("Error Connecting to cluster")
	}

	children, _, err := conn.Children(path)
	if err != nil {
		log.Error("Error fetching children")
	}
	return children

}
