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
	log.Info(zoodashConfig)
}
