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

	err := board.ChangeBoard(changes)
	if err != nil {
		fmt.Println(err)

	}
	board.Print()
}
