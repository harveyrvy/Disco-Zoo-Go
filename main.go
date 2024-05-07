package main

import (
	"fmt"

	farm "github.com/discozoo/Animal/FarmAnimals"
	boardpkg "github.com/discozoo/Board"
)

func main() {

	board := boardpkg.New()
	changes := []boardpkg.BoardChange{
		boardpkg.NewBoardChange(0, 0, farm.Rabbit),
		boardpkg.NewBoardChange(0, 1, farm.Sheep),
		boardpkg.NewBoardChange(4, 0, farm.Chicken),
	}

	for _, c := range changes {
		err := board.ChangeBoard(c)
		if err != nil {
			fmt.Println(err)

		}
		//board.Print()
	}

	board2 := boardpkg.New()
	matrix := board2.GetMatrix()
	for i := range matrix {
		for j := range matrix[i] {
			err := board2.ChangeBoard(boardpkg.NewBoardChange(i, j, farm.Rabbit))

			if err != nil {
				fmt.Println(err)
				continue
			}
			board2.Print()
			for i := range matrix {
				for j := range matrix[i] {
					board2.IncCounts(boardpkg.NewBoardChange(i, j, farm.Sheep))
				}
			}
			board2.ClearAnimals()
		}
	}
	board2.Print()

}
