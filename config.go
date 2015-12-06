package main

import (
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var zoodashConfig Config

// Config is a basic configuration in a yaml file
type Config struct {
	Path     string
	Adresses map[string]string `yaml:"servers,omitempty"`
}

func init() {
	zoodashConfig = Config{Path: "/etc/zoodash/config.yml"}

	if source, err := ioutil.ReadFile(zoodashConfig.Path); err == nil {
		err = yaml.Unmarshal([]byte(source), &zoodashConfig)
		if err != nil {
			log.Panic(err)
		}
	}
	log.Debug(zoodashConfig)
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
