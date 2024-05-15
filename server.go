package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"regexp"

	calcs "github.com/discozoo/Calcs"
	region "github.com/discozoo/Region"
)

type HomePageVariables struct {
	Grid        string
	BoardsCount int
}

type RegionPageVariables struct {
	RegionName  string
	AnimalNames []string
	Path        string
}

var regionMap = region.GetRegionMap()

func HomePage(w http.ResponseWriter, r *http.Request) {
	boardsCount, prc := calcs.Calculation()
	HomePageVars := HomePageVariables{
		Grid:        prc.String(),
		BoardsCount: boardsCount,
	}
	t, err := template.ParseFiles("homepage.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("template execution error: ", err)
	}
}

func RegionPage() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		region := filepath.Base(r.URL.Path)
		RegionPageVars := RegionPageVariables{
			Path:        r.URL.Path,
			RegionName:  region,
			AnimalNames: regionMap[region]} //TODO: make this get from a map with key region
		t, err := template.ParseFiles("regionpage.html")
		if err != nil {
			log.Print("template parsing error: ", err)
		}
		err = t.Execute(w, RegionPageVars)
		if err != nil {
			log.Print("template execution error: ", err)
		}
	}
}

// Regular expression handler
type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHandler struct {
	routes []*route
}

func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}

func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	//handler := &RegexpHandler{}
	//handler.HandleFunc(regexp.MustCompile("/"), HomePage)
	//handler.HandleFunc(regexp.MustCompile("[a-z]*$"), HomePage)
	http.HandleFunc("/home/", HomePage)
	http.HandleFunc("/", RegionPage())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
