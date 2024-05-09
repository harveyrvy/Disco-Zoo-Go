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
