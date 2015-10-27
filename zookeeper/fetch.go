package zookeeper

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/prfalken/zoodash/logger"
)

// interval of zookeeper stat checks
const ZKCheckInterval = 1 * time.Second

func RunStatsFetcher(zookeeper *Zookeeper, ZKserver string) {
	for {
		timeout := 3 * time.Second
		conn, err := net.DialTimeout("tcp", ZKserver, timeout)
		if err != nil {
			logger.Warning.Println("Could not tcp dial Zookeeper")
		}
		out := get4LettersFromZookeeper(conn, "stat")
		ParseStatsOutput(out, &zookeeper.Statistics)
		conn.Close()

		// To be refactored - How to keep the same connection over time ?
		conn, err = net.DialTimeout("tcp", ZKserver, timeout)
		defer conn.Close()
		if err != nil {
			logger.Warning.Println("Could not tcp dial Zookeeper")
		}
		out = get4LettersFromZookeeper(conn, "envi")
		ParseEnvOutput(out, &zookeeper.Environment)
		conn.Close()

		time.Sleep(ZKCheckInterval)
	}
}

func get4LettersFromZookeeper(conn net.Conn, cmd string) string {
	fmt.Fprintf(conn, cmd+"\r")
	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 1024)
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				logger.Warning.Println("Could not finish data fetch from ZK:", err)
			}
			break
		}
		buf = append(buf, tmp[:n]...)

	}
	return string(buf)

}
