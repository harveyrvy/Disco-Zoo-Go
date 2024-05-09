package main

import (
	"fmt"

	animal "github.com/discozoo/Animal"
	farm "github.com/discozoo/Animal/FarmAnimals"
	boardpkg "github.com/discozoo/Board"
)

func main() {

	var animals = []animal.Animal{farm.Rabbit, farm.Sheep, farm.Chicken}
	allBoards := generateAllBoards(animals)

	//printAllBoards(allBoards)

	fmt.Printf("There are %d boards \n", len(allBoards))

	prc := boardpkg.PercCount{}
	prcMatrix := prc.GetMatrix()
	for i := range prcMatrix {
		for j := range prcMatrix[i] {
			for _, b := range allBoards {
				bMatrix := b.GetMatrix()
				if bMatrix[i][j].GetState() {
					prc.IncValue(i, j)
				}
			}
			// convert to percentiles calculation
			prc.ScaleValue(i, j, 100/float64(len(allBoards)))
		}

	}
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
					fmt.Println(err)
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
