package main

import (
	"fmt"

	animal "github.com/discozoo/Animal"
	farm "github.com/discozoo/Animal/FarmAnimals"
	boardpkg "github.com/discozoo/Board"
)

var placementErrorLog = []string{}

func main() {
	var animals = []animal.Animal{farm.Rabbit, farm.Sheep}
	allBoards := generateAllBoards(animals)

	//printAllBoards(allBoards)
	//fmt.Print(placementErrorLog)

	fmt.Printf("There are %d boards \n", len(allBoards))

	prc := boardpkg.CalculatePercentages(allBoards)
	fmt.Println(prc)
}

func generateAllBoards(animals []animal.Animal) [](boardpkg.Board) {

	board := boardpkg.New()
	boards := []boardpkg.Board{board}
	animalCount := len(animals)

	for i := range animalCount {
		boards = allAnimalPositionsForBoard(animals[i], boards)
	}

	return boards
}

// Given list of existing boards, and an animal, generate list of all boards adding that animal to given boards.
func allAnimalPositionsForBoard(animal animal.Animal, boards []boardpkg.Board) []boardpkg.Board {
	output := []boardpkg.Board{}
	for _, b := range boards {
		matrix := b.GetMatrix()
		bCopy := b
		for i := range matrix {
			for j := range matrix {
				err := bCopy.ChangeBoard(boardpkg.NewBoardChange(i, j, animal))

				if err != nil {
					placementErrorLog = append(placementErrorLog, fmt.Sprintf("%s\n", err))
					continue
				}
				// add new board
				output = append(output, bCopy)
				// reset to board without previous animal added
				bCopy = b
			}
		}
	}

	return output
}

func printAllBoards(boards []boardpkg.Board) {
	for _, b := range boards {
		fmt.Println(b)
	}
}
