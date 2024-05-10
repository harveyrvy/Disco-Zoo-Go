package main

import (
	"fmt"
	"net/http"

	calcs "github.com/discozoo/Calcs"
)

func printBoard(w http.ResponseWriter, r *http.Request) {
	boardsCount, prc := calcs.Calculation()
	fmt.Fprintf(w, "There are %d boards \n", boardsCount)
	fmt.Fprintf(w, "%v", prc)
}

func main() {
	http.HandleFunc("/", printBoard)
	http.ListenAndServe(":8080", nil)
}
