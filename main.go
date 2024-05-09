package main

import (
	"fmt"

	animal "github.com/discozoo/Animal"
	farm "github.com/discozoo/Animal/FarmAnimals"
	boardpkg "github.com/discozoo/Board"
)

func main() {

	boardsCount := 0
	board := boardpkg.New()
	matrix := board.GetMatrix()

	var animals = []animal.Animal{}

	for i := range matrix {
		for j := range matrix[i] {
			err := board.ChangeBoard(boardpkg.NewBoardChange(i, j, farm.Rabbit))

			if err != nil {
				fmt.Println(err)
				continue
			}
			for i := range matrix {
				for j := range matrix[i] {
					n := board.IncCounts(boardpkg.NewBoardChange(i, j, farm.Sheep))
					boardsCount = boardsCount + n
				}
			}
			board.ClearAnimals()
		}
	}
	fmt.Println(board)
	fmt.Println(boardsCount)
	fmt.Println(board.GetCounts().ConvertCounts(boardsCount))

	rabbitBoards := allAnimalPositionsForBoard(farm.Rabbit, []boardpkg.Board{boardpkg.New()})
	sheepBoards := allAnimalPositionsForBoard(farm.Sheep, rabbitBoards)

	for _, b := range sheepBoards {
		fmt.Println(b)
	}

	allBoards := generateAllBoards(animals)
	for _, b := range allBoards {
		fmt.Println(b)
	}

}

func generateAllBoards(animals []animal.Animal) [](boardpkg.Board) {

	//animalCount := len(animals)
	board := boardpkg.New()
	boards := []boardpkg.Board{board}

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
