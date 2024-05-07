package Board

import (
	"fmt"

	animal "github.com/discozoo/Animal"
)

type Board struct {
	matrix [5][5]Tile
	counts [5][5]int
}

func New() Board {
	matrix := [5][5]Tile{}
	counts := [5][5]int{}

	t := NewBlankTile()
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = t
		}
	}

	return Board{matrix, counts}
}

func boardChangePossible(b Board, startX, startY int, animal animal.Animal) error {
	if startX < 0 || startY < 0 || startX > 4 || startY > 4 {
		return fmt.Errorf("start tile outside of board")
	}
	for _, v := range animal.GetTiles() {
		if startX+v[0] < 0 || startY+v[1] < 0 || startX+v[0] > 4 || startY+v[1] > 4 {
			return fmt.Errorf("animal will be placed outside of board")
		}
		if b.matrix[startX+v[0]][startY+v[1]].state {
			return fmt.Errorf("replacing over another animal")
		}
	}
	return nil
}

func (b *Board) ChangeBoard(changes []BoardChange) error {
	for _, c := range changes {
		fmt.Println(c)
		err := boardChangePossible(*b, c.startX, c.startY, c.animal)
		if err != nil {
			return fmt.Errorf("board change %v wasn't possible: %s", c, err)
		}
		for _, v := range c.animal.GetTiles() {
			b.matrix[c.startX+v[0]][c.startY+v[1]] = NewTile(c.animal)
			b.counts[c.startX+v[0]][c.startY+v[1]]++
		}
	}
	return nil
}

func (b *Board) String() string {
	str := ""
	for i := range b.matrix {
		str = str + "||  "
		for j := range b.matrix[i] {
			str = str + fmt.Sprintf("%8s | %-8d  ||  ", b.matrix[i][j].animal.String(), b.counts[i][j])
		}
		str = str + "\n"
	}
	return str
}

func (b *Board) Print() {
	fmt.Print(b.String())

}
