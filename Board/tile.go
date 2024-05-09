package Board

import (
	"fmt"

	animal "github.com/discozoo/Animal"
)

type Tile struct {
	state  bool
	animal animal.Animal
}

func NewBlankTile() Tile {
	tile := Tile{false, animal.NewDefault()}
	return tile
}

func NewTile(a animal.Animal) Tile {
	tile := Tile{true, a}
	return tile
}

func (t *Tile) SetState(b bool) {
	t.state = b
}

func (t *Tile) GetState() bool {
	return t.state
}

func (t *Tile) Print() {
	if !t.state {
		fmt.Printf("Empty")
	} else {
		fmt.Print(t.animal)
	}
}
