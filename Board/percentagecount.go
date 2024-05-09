package Board

import "fmt"

type PercCount struct {
	matrix [5][5]float64
}

func (p *PercCount) GetMatrix() [5][5]float64 {
	return p.matrix
}

func (p *PercCount) SetValue(i, j int, val float64) {
	p.matrix[i][j] = val
}

func (p *PercCount) IncValue(i, j int) {
	p.matrix[i][j]++
}

func (p *PercCount) ScaleValue(i, j int, mult float64) {
	p.matrix[i][j] = p.matrix[i][j] * mult
}

func CalculatePercentages(boards []Board) PercCount {
	matrix := [5][5]float64{}
	p := PercCount{matrix}
	for i := range matrix {
		for j := range matrix[i] {
			for _, b := range boards {
				if b.matrix[i][j].GetState() {
					p.IncValue(i, j)
				}
			}
			// convert to percentiles calculation
			p.ScaleValue(i, j, 100/float64(len(boards)))
		}
	}
	return p
}

func (p PercCount) String() string {
	str := ""
	for i := range p.matrix {
		str = str + "||  "
		for j := range p.matrix[i] {
			str = str + fmt.Sprintf("  %.2f%%  ||", p.matrix[i][j])
		}
		str = str + "\n"
	}
	return str
}
