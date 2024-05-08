package Board

import (
	"fmt"

	animal "github.com/discozoo/Animal"
)

type BoardChange struct {
	startX, startY int
	animal         animal.Animal
}

func NewBoardChange(startX, startY int, a animal.Animal) BoardChange {

	return BoardChange{startX, startY, a}
}

func (c BoardChange) String() string {
	return fmt.Sprintf("[%s at %d, %d]", c.animal.GetName(), c.startX, c.startY)
}
