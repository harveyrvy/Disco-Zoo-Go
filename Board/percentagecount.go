package Board

type percCount struct {
	matrix [5][5]float64
}

func (p *percCount) setValue(i, j int, val float64) {
	p.matrix[i][j] = val
}
