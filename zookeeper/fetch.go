package zookeeper

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/prfalken/zoodash/logger"
)

// interval of zookeeper stat checks
const ZKCheckInterval = 5 * time.Second
const zkDialTimeout = 10 * time.Second

func RunFetcher(zookeeper *Zookeeper, ZKserver string) {

	commands := map[string]OutputParser{
		"stat": &zookeeper.Statistics,
		"envi": &zookeeper.Environment,
	}

	for {

		for command, zkPageType := range commands {
			conn, err := net.DialTimeout("tcp", ZKserver, zkDialTimeout)
			if err != nil {
				logger.Warning.Println("Could not tcp dial Zookeeper")
			}
			out := get4LettersFromZookeeper(conn, command)
			zkPageType.ParseOutput(out)
			conn.Close()

		}
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
