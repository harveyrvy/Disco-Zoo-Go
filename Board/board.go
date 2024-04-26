package board

type Board struct {
	matrix [5][5]Tile
	counts [5][5]int
}

func NewDefaultBoard() Board {
	matrix := [5][5]Tile{}
	counts := [5][5]int{}

	return Board{matrix, counts}

}
