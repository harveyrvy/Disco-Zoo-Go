package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	boardpkg "github.com/discozoo/Board"
	calcs "github.com/discozoo/Calcs"
	region "github.com/discozoo/Region"
)

type HomePageVariables struct {
	Grid        string
	BoardsCount int
	Boards      []boardpkg.Board
}

type RegionPageVariables struct {
	RegionName  string
	AnimalNames []string
	Path        string
}

var regionMap = region.GetRegionMap()

func HomePage(w http.ResponseWriter, r *http.Request) {
	boardsCount, prc, allBoards:= calcs.Calculation()
	Boards := allBoards
	HomePageVars := HomePageVariables{
		Grid:        prc.String(),
		Boards:      Boards,
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

func CalculateHandler(w http.ResponseWriter, r *http.Request) {}


func main() {
	http.HandleFunc("/calculate/". CalculateHandler)
	http.HandleFunc("/home/", HomePage)
	http.HandleFunc("/", RegionPage())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
