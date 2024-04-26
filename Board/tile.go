package board

import (
	"fmt"

	animal "github.com/discozoo/Animal"
)

type Tile struct {
	state  bool
	animal animal.Animal
}

func (t Tile) Print() {
	if !t.state {
		fmt.Printf("Empty")
	} else {
		fmt.Print(t.animal)
	}
}
