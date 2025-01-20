package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

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

type Request struct {
	Animals []string `json:animals`
}

type AnimalResponse struct {
	Animals []string `json:response`
}

var regionMap = region.GetRegionMap()

func HomePage(w http.ResponseWriter, r *http.Request) {
	boardsCount, prc, allBoards := calcs.Calculation()
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

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	animalName := r.URL.Query().Get("animal")

	response := map[string]string{"result": animalName}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := req.Animals
	fmt.Println(result)

	// Return the result in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AnimalResponse{Animals: result})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/postcalculate/", responseHandler).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/calculate/", CalculateHandler)
	r.HandleFunc("/home/", HomePage)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	http.Handle("/", handlers.CORS(originsOk, headersOk, methodsOk)(r))
	//http.HandleFunc("/", RegionPage())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
