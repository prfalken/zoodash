package main

import (
	"bytes"
	"io/ioutil"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/samuel/go-zookeeper/zk"
)

func loadTemplate(filename string) string {
	filename = "templates/" + filename
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error("Could not load template " + filename)
	}
	return string(body)
}

func index(stats []zk.ServerStats) string {
	tmpl := loadTemplate("index.html")
	t, err := template.New("index").Parse(tmpl)
	if err != nil {
		log.Error("could not parse template " + tmpl)
	}
	buf := new(bytes.Buffer)
	log.Info(stats)
	err = t.Execute(buf, stats)
	rendered := buf.String()
	return rendered
}
