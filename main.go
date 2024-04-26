package main

import (
	"fmt"

	animal "github.com/discozoo/Animal"
	board "github.com/discozoo/Board"
)

func main() {

	board := board.New()
	board.Print()
	err := board.ChangeBoard(0, 0, animal.Rabbit)
	if err != nil {
		fmt.Println(err)
	}
	board.Print()

}
