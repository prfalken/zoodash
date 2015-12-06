package main

import (
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/samuel/go-zookeeper/zk"
)

type Node map[string]string

func getNodeChildren(path string) ([]Node, error) {

	var nodes []Node

	ipPorts := getZKAddressesFromConfig()
	children := []string{}
	conn, _, err := zk.Connect(ipPorts, 2*time.Second)
	if err != nil {
		log.Error("Error Connecting to cluster")
		return nil, err
	}

	// remove trailing slash, not supported by zk
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}

	children, _, err = conn.Children(path)
	if err != nil {
		log.Error("Error fetching children")
		return nodes, err
	}

	for _, c := range children {
		node := Node{
			"title": c,
			"key":   path + "/" + c,
			"lazy":  "true",
		}
		log.Debug(node)
		nodes = append(nodes, node)
	}

	return nodes, nil

}
