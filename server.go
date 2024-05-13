package main

import (
	"html/template"
	"log"
	"net/http"

	farm "github.com/discozoo/Animal/Farm"
	outback "github.com/discozoo/Animal/Outback"
	calcs "github.com/discozoo/Calcs"
)

var RegionAnimals = map[string]([]string){
	"farm":    farm.GetAnimalNames(),
	"outback": outback.GetAnimalNames(),
}

type HomePageVariables struct {
	Grid        string
	BoardsCount int
}

type RegionPageVariables struct {
	RegionName string
	AnimalList []string
}

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

func RegionPage(region string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		animalNames := RegionAnimals[region]
		FarmPageVars := RegionPageVariables{
			RegionName: region,
			AnimalList: animalNames}
		t, err := template.ParseFiles("regionpage.html")
		if err != nil {
			log.Print("template parsing error: ", err)
		}
		err = t.Execute(w, FarmPageVars)
		if err != nil {
			log.Print("template execution error: ", err)
		}
	}
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/farm/", RegionPage("farm"))
	http.HandleFunc("/outback/", RegionPage("outback"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
