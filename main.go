package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/prfalken/zoodash/logger"
	zk "github.com/prfalken/zoodash/zookeeper"
)

func main() {
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	ZKservers := []string{
		"192.168.33.10:2181",
		"192.168.33.10:2182",
	}
	var zookeepers []*zk.Zookeeper
	for _, server := range ZKservers {
		zookeeper := &zk.Zookeeper{}
		zookeepers = append(zookeepers, zookeeper)
		go zk.RunFetcher(zookeeper, server)
	}

	go func() {

		for {
			for _, zook := range zookeepers {
				fmt.Println(zook)
			}

			time.Sleep(1 * time.Second)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, index(&zookeepers))
	})
	http.ListenAndServe(":8080", nil)

}
