package main

import (
	"fmt"

	farm "github.com/discozoo/Animal/FarmAnimals"
	boardpkg "github.com/discozoo/Board"
)

func main() {

	boardsCount := 0

	board := boardpkg.New()
	matrix := board.GetMatrix()
	for i := range matrix {
		for j := range matrix[i] {
			err := board.ChangeBoard(boardpkg.NewBoardChange(i, j, farm.Rabbit))

			if err != nil {
				fmt.Println(err)
				continue
			}
			board.Print()
			for i := range matrix {
				for j := range matrix[i] {
					n := board.IncCounts(boardpkg.NewBoardChange(i, j, farm.Sheep))
					boardsCount = boardsCount + n
				}
			}
			board.ClearAnimals()
		}
	}
	board.Print()
	fmt.Println(boardsCount)
	fmt.Println(board.GetCounts().ConvertCounts(boardsCount))

}
