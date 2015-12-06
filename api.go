package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"github.com/samuel/go-zookeeper/zk"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, buildIndexPage())
}

func (s *StatsStorage) statsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stats := s.getStats()
	fmt.Fprintf(w, buildStatsPage(stats))
}

func apiHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var parent string
	parentParam := r.URL.Query()["parent"]
	if len(parentParam) > 0 {
		parent = parentParam[0]
	} else {
		parent = "/"
	}
	log.Error(parent)
	w.Header().Set("Content-Type", "application/json")
	nodes, err := getNodeChildren(parent)

	if err != nil {
		fmt.Fprint(w, "Error")
	}
	if err := json.NewEncoder(w).Encode(nodes); err != nil {
		fmt.Fprint(w, "Error")
	}

}

func browseHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, buildBrowsePage())
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

func buildIndexPage() string {
	return buildPage("index.html", nil)
}

func buildStatsPage(stats map[string]zk.ServerStats) string {
	return buildPage("stats.html", stats)
}

func buildBrowsePage() string {
	return buildPage("browse.html", nil)
}
