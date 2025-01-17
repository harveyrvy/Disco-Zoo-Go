package calculations

import (
	"fmt"

	animal "github.com/discozoo/Animal"
	boardpkg "github.com/discozoo/Board"
	farm "github.com/discozoo/Region/Farm"
)

var placementErrorLog = []string{}

func Calculation() (boardCount int, prc boardpkg.PercCount, allBoards [](boardpkg.Board)) {
	var animals = []animal.Animal{farm.Sheep, farm.Rabbit}
	allBoards = generateAllBoards(animals)

	//printAllBoards(allBoards)
	//fmt.Print(placementErrorLog)

	//fmt.Printf("There are %d boards \n", len(allBoards))

	prc = boardpkg.CalculatePercentages(allBoards)
	return len(allBoards), prc, allBoards
}

func Calculation2(animals []animal.Animal) (boardCount int, prc boardpkg.PercCount, allBoards [](boardpkg.Board)) {
	allBoards = generateAllBoards(animals)
	prc = boardpkg.CalculatePercentages(allBoards)
	return len(allBoards), prc, allBoards
}

// Given list of animals, generate all boards with those animals
func generateAllBoards(animals []animal.Animal) [](boardpkg.Board) {

	board := boardpkg.New()
	boards := []boardpkg.Board{board}
	animalCount := len(animals)

	for i := range animalCount {
		boards = allAnimalPositionsForBoard(animals[i], boards)
	}

	return boards
}

// Given list of existing boards, and an animal, generate list of all boards which add that animal to given boards.
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
				// add new board to output list
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
