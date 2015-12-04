package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"github.com/samuel/go-zookeeper/zk"
)

func (s *StatsStorage) IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stats := s.getStats()
	fmt.Fprintf(w, buildStatsPage(stats))
}

func browseHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := p.ByName("filepath")
	fmt.Fprintf(w, buildBrowsePage(path))
}

func loadTemplate(filename string) string {
	filename = "templates/" + filename
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error("Could not load template " + filename)
	}
	return string(body)
}

func buildPage(pageName string, context interface{}) string {
	tmpl := loadTemplate(pageName)
	t, err := template.New(pageName).Parse(tmpl)
	if err != nil {
		log.Error("could not parse template " + tmpl)
	}
	buf := new(bytes.Buffer)
	log.Debug(context)
	err = t.Execute(buf, context)
	rendered := buf.String()
	return rendered
}

func buildStatsPage(stats map[string]zk.ServerStats) string {
	return buildPage("index.html", stats)
}

func buildBrowsePage(path string) string {
	children := ZKBrowse(path)
	return buildPage("browse.html", children)
}
