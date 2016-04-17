package matrix

// TODO: make more complex
func (m *Matrix) Insert(row ...[]float64) *Matrix {
	for index := range row {
		if m.cols == 0 {
			m.cols = len(row[index])
		} else {
			if len(row[index]) != m.cols {
				panic("Vertical dimensions mismatch")
			}
		}
		m.Table = append(m.Table, row[index])
		m.rows = m.rows + 1
	}
	return m
}