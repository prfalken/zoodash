package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	ZKservers := []string{
		"192.168.33.10:2181",
		"192.168.33.10:2182",
		"192.168.33.10:2183",
	}
	var stats []zk.ServerStats
	go func() {
		for {

			serverStats, ok := zk.FLWSrvr(ZKservers, 3*time.Second)
			buff := []zk.ServerStats{}
			for _, s := range serverStats {
				buff = append(buff, *s)
			}
			stats = buff
			if !ok {
				Error.Println("Failed to fetch stats on one or more servers")
			}

			time.Sleep(2 * time.Second)
		}

	}()

	// go func() {
	// 	for {
	// 		for _, s := range stats {
	// 			Info.Println(s)
	// 		}
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, index(stats))
	})
	http.ListenAndServe(":8080", nil)

}
