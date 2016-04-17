package matrix

import (
	"strconv"
	log "github.com/Sirupsen/logrus"
	"github.com/JonathanRosado/Shako/linalg/matrix/data"
)

func (m *Matrix) vectorIndex(pos int) float64{
	if pos >= (m.data.Rows*m.data.Cols) {
		log.WithFields(log.Fields{
			"rowsXcols": (m.data.Rows*m.data.Cols),
			"vectorIndex": pos,
		}).Warn("vectorIndex out of bounds")
		return -1.0
	}

	// i(vector_index) = vector_index/number_of_columns
	i := pos/m.data.Cols

	// j(vector_index) = vector_index - (number_of_columns*i)
	j := pos - (m.data.Cols * i)

	return m.data.Get(i,j)
}

func (m *Matrix) matrixIndex(pi int, pj int) float64{
	return m.data.Get(pi, pj)
}

func (m *Matrix) i_s_index(i int, j string) *Matrix {
	switch len(j) {
	case 1:
		if j == ":" {
			matrix := &Matrix{}
			matrix.data = *(&data.Data{}).Create([][]float64{
				m.data.GetRow(i),
			})
			return matrix
		}
		if jVal, err := strconv.Atoi(j); err == nil {
			matrix := &Matrix{}
			matrix.data = *(&data.Data{}).Create([][]float64{
				[]float64{m.Elem(i, jVal)},
			})
			return matrix
		} else {
			panic("Index operation not supported")
		}
	}

	matrix := &Matrix{}
	matrix.data = *(&data.Data{}).Create([][]float64{
		[]float64{0},
	})
	return matrix
}

func (m *Matrix) s_i_index(i string, j int) *Matrix {
	switch len(i) {
	case 1:
		if i == ":" {
			matrix := &Matrix{}
			matrix.data = *(&data.Data{}).Create([][]float64{
				m.data.GetColumn(j),
			}).Transpose()
			return matrix
		}
		if iVal, err := strconv.Atoi(i); err == nil {
			matrix := &Matrix{}
			matrix.data = *(&data.Data{}).Create([][]float64{
				[]float64{m.Elem(iVal, iVal)},
			})
			return matrix
		} else {
			panic("Index operation not supported")
		}
	}

	matrix := &Matrix{}
	matrix.data = *(&data.Data{}).Create([][]float64{
		[]float64{0},
	})
	return matrix
}

func (m *Matrix) i_i_index(i int, j int) *Matrix {
	matrix := &Matrix{}
	matrix.data = *(&data.Data{}).Create([][]float64{
		[]float64{m.matrixIndex(i, j)},
	})
	return matrix
}

func (m *Matrix) s_s_index(i string, j string) *Matrix {
	type IndexType int

	const (
		col IndexType = iota
		intCol
		integer
	)

	ijTypes := []IndexType{}

	if _, err := strconv.Atoi(i); err == nil {
		ijTypes = append(ijTypes, integer)
	} else if i == ":" {
		ijTypes = append(ijTypes, col)
	} else {
		panic("Index operation not supported")
	}

	if _, err := strconv.Atoi(j); err == nil {
		ijTypes = append(ijTypes, integer)
	} else if j == ":" {
		ijTypes = append(ijTypes, col)
	} else {
		panic("Index operation not supported")
	}

	if ijTypes[0] == col && ijTypes[1] == col {
		return m
	} else if ijTypes[0] == integer && ijTypes[1] == integer {
		pi, err := strconv.Atoi(i)
		if err != nil{
			panic("Atoi failed")
		}
		pj, err1 := strconv.Atoi(j)
		if err1 != nil{
			panic("Atoi failed")
		}
		return m.i_i_index(pi, pj)
	} else if ijTypes[0] == integer && ijTypes[1] == col {
		pi, err := strconv.Atoi(i)
		if err != nil{
			panic("Atoi failed")
		}
		return m.i_s_index(pi, j)
	} else if ijTypes[0] == col && ijTypes[1] == integer {
		pj, err := strconv.Atoi(j)
		if err != nil{
			panic("Atoi failed")
		}
		return m.s_i_index(i, pj)
	}

	matrix := &Matrix{}
	matrix.data = *(&data.Data{}).Create([][]float64{
		[]float64{0},
	})
	return matrix
}

//func (m *Matrix) rowVector(pos int) *Matrix {
//
//}

func (m *Matrix) columnVector() *Matrix {
	m.data.ColumnVector()
	return m
}