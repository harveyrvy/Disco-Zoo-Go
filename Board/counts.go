package Board

type countsMatrix struct {
	counts [5][5]int
}

func (c *countsMatrix) getValue(i, j int) int {
	return c.counts[i][j]
}

func (c *countsMatrix) ConvertCounts(boards int) PercCount {
	prc := PercCount{}
	for i := range c.counts {
		for j := range c.counts[i] {
			prc.SetValue(i, j, 100*float64(c.counts[i][j])/float64(boards))

		}
	}
	return prc
}

func (c *countsMatrix) IncrementCount(i, j int) {
	c.counts[i][j]++
}
