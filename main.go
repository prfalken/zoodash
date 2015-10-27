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
	var zkstats []*zk.Stat
	var zkenvs []*zk.Env
	for _, server := range ZKservers {
		zkstat := &zk.Stat{}
		zkenv := &zk.Env{}
		zkstats = append(zkstats, zkstat)
		zkenvs = append(zkenvs, zkenv)

		go zk.RunStatsFetcher(zkstat, zkenv, server)
	}

	go func() {

		for {
			for _, stat := range zkstats {
				fmt.Println(stat)
			}
			for _, env := range zkenvs {
				fmt.Println(env)
			}

			time.Sleep(1 * time.Second)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, index(&zkstats, &zkenvs))
	})
	http.ListenAndServe(":8080", nil)

}
