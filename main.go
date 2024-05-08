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

	//var animals = []animal.Animal

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
	fmt.Println(board.ConvertCounts(boardsCount))
	animals := []animal.Animal{farm.Rabbit, farm.Sheep}
	boards := generateAllBoards(animals)
	for _, b := range boards {
		fmt.Println(b)
	}

}

func generateAllBoards(animals []animal.Animal) [](boardpkg.Board) {
	boards := [](boardpkg.Board){}
	//animalCount := len(animals)
	board := boardpkg.New()
	matrix := board.GetMatrix()
	for i := range matrix {
		for j := range matrix {
			err := board.ChangeBoard(boardpkg.NewBoardChange(i, j, animals[0])) // TODO: change to be length of animals list

			if err != nil {
				fmt.Println(err)
				continue
			}
			board1 := board
			for i := range matrix {
				for j := range matrix[i] {
					err := board.ChangeBoard(boardpkg.NewBoardChange(i, j, animals[1]))
					if err != nil {
						fmt.Println(err)
						continue
					}
					boards = append(boards, board)
					board = board1

				}
			}
			board = boardpkg.New()
			board1 = boardpkg.New()
		}
	}

	return boards
}
