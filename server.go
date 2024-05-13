package main

import (
	"html/template"
	"log"
	"net/http"

	farm "github.com/discozoo/Animal/Farm"
	calcs "github.com/discozoo/Calcs"
)

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

func FarmPage(w http.ResponseWriter, r *http.Request) {
	animalNames := farm.GetAnimalNames()
	FarmPageVars := RegionPageVariables{
		RegionName: "farm",
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

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/farm/", FarmPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
