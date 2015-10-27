package main

import (
	"bytes"
	"io/ioutil"
	"text/template"

	"github.com/prfalken/zoodash/logger"
	zk "github.com/prfalken/zoodash/zookeeper"
)

func loadTemplate(filename string) string {
	filename = "templates/" + filename
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Error.Println("Could not load template " + filename)
	}
	return string(body)
}

func index(stats *[]*zk.Stat, envs *[]*zk.Env) string {
	tmpl := loadTemplate("index.html")
	t, err := template.New("index").Parse(tmpl)
	if err != nil {
		logger.Error.Println("could not parse template " + tmpl)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, stats)
	rendered := buf.String()
	return rendered
}
