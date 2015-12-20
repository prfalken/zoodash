package main

import (
	"flag"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var zoodashConfig Config

// Config is a basic configuration in a yaml file
type Config struct {
	Path          string
	Adresses      map[string]string `yaml:"servers,omitempty"`
	ListenAddress string
	ListenPort    string
}

func init() {
	zoodashConfig = Config{Path: "/etc/zoodash/config.yml"}

	if source, err := ioutil.ReadFile(zoodashConfig.Path); err == nil {
		err = yaml.Unmarshal([]byte(source), &zoodashConfig)
		if err != nil {
			log.Panic(err)
		}
	}

	overrideWithCommandLine(&zoodashConfig)

	log.Debug(zoodashConfig)
}

func overrideWithCommandLine(config *Config) {
	lAddress := flag.String("listen", "127.0.0.1", "IP address to listen to")
	lPort := flag.String("port", "8080", "TCP port to listen to")
	flag.Parse()
	config.ListenAddress = *lAddress
	config.ListenPort = *lPort

}

func getZKAddressesFromConfig() []string {

	ipPorts := []string{}
	for _, ipPort := range zoodashConfig.Adresses {
		ipPorts = append(ipPorts, ipPort)
	}
	log.Debug(ipPorts)
	return ipPorts
}

func getZKHostsFromConfig() []string {
	hosts := []string{}

	for host := range zoodashConfig.Adresses {
		hosts = append(hosts, host)
	}
	return hosts
}
