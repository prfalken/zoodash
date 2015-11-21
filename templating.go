package main

import (
	"bytes"
	"io/ioutil"
	"text/template"

	"github.com/samuel/go-zookeeper/zk"
)

func loadTemplate(filename string) string {
	filename = "templates/" + filename
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		Error.Println("Could not load template " + filename)
	}
	return string(body)
}

func index(stats []zk.ServerStats) string {
	tmpl := loadTemplate("index.html")
	t, err := template.New("index").Parse(tmpl)
	if err != nil {
		Error.Println("could not parse template " + tmpl)
	}
	buf := new(bytes.Buffer)
	Info.Println(stats)
	err = t.Execute(buf, stats)
	rendered := buf.String()
	return rendered
}
