package matrix

// TODO: make more complex
func (m *Matrix) Insert(row ...[]float64) *Matrix {
	for index := range row {
		if m.verticalDim == 0 {
			m.verticalDim = len(row[index])
		} else {
			if len(row[index]) != m.verticalDim {
				panic("Vertical dimensions mismatch")
			}
		}
		m.Table = append(m.Table, row[index])
		m.horizontalDim = m.horizontalDim + 1
	}
	return m
}