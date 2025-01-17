package Utils

type TilesList = [][2]int

func NewTilesList() TilesList {
	l := make([][2]int, 0)
	return l
}

func New(x,y []int) (TilesList, err){
	l := make([][2]int, 0)
	if len(x) != len(y) {
		e:= "List of coordinates" 
	}
	
	
}
