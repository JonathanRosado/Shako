package matrix

// TODO: make more complex
func (m *Matrix) Create(rows ...[]float64) *Matrix {
	matrix := &Matrix{}
	matrix.data.Create([][]float64(rows))

	return matrix
}