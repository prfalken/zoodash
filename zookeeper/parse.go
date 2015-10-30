package zookeeper

import (
	"regexp"
	"strconv"

	"github.com/prfalken/zoodash/logger"
)

type Client struct {
	ip       string
	port     int
	queued   int
	received int
	sent     int
}

type Stat struct {
	Clients     string
	Latency     string
	Received    int
	Sent        int
	Connections int
	Outstanding int
	Zxid        string
	Mode        string
	NodeCount   int
}

type Env struct {
	Hostname    string
	JavaVersion string
	OSname      string
	OSarch      string
	OSversion   string
}

type Zookeeper struct {
	Environment Env
	Statistics  Stat
}

type OutputParser interface {
	ParseOutput(string)
}

func catch(reg, text string) string {
	re := regexp.MustCompile(reg)
	return re.FindStringSubmatch(text)[1]

}

func (stat *Stat) ParseOutput(zkOutput string) {
	stat.Latency = catch("Latency min/avg/max: (.*)", zkOutput)
	stat.Zxid = catch("Zxid: (.*)", zkOutput)
	stat.Mode = catch("Mode: (.*)", zkOutput)

	received, err := strconv.Atoi(catch("Received: (.*)", zkOutput))
	if err != nil {
		logger.Warning.Println("Couldnt parse Received field")
	}
	sent, err := strconv.Atoi(catch("Sent: (.*)", zkOutput))
	if err != nil {
		logger.Warning.Println("Couldnt parse Sent field")
	}
	connections, err := strconv.Atoi(catch("Connections: (.*)", zkOutput))
	if err != nil {
		logger.Warning.Println("Couldnt parse Connections field")
	}
	outstanding, err := strconv.Atoi(catch("Outstanding: (.*)", zkOutput))
	if err != nil {
		logger.Warning.Println("Couldnt parse Outstanding field")
	}
	nodeCount, err := strconv.Atoi(catch("Node count: (.*)", zkOutput))
	if err != nil {
		logger.Warning.Println("Couldnt parse Node Count field")
	}

	stat.Received = received
	stat.Sent = sent
	stat.Connections = connections
	stat.Outstanding = outstanding
	stat.NodeCount = nodeCount
}

func (env *Env) ParseOutput(zkOutput string) {
	env.Hostname = catch("host.name=(.*)", zkOutput)
}
