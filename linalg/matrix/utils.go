package matrix

import "strconv"

func (m *Matrix) vectorIndex(pos int) float64{
	var count int = 0
	var elem float64 = -1.0
	for i, item := range m.Table {
		for j := range item {
			if count == pos {
				elem = m.Table[i][j]
			}
			count = count + 1
		}
	}
	return elem
}

func (m *Matrix) matrixIndex(pi int, pj int) float64{
	return m.Table[pi][pj]
}

func (m *Matrix) i_s_index(i int, j string) *Matrix {
	switch len(j) {
	case 1:
		if j == ":" {
			return (&Matrix{}).Insert(
				m.Table[i],
			)
		}
		if jVal, err := strconv.Atoi(j); err == nil {
			return (&Matrix{}).Insert(
				Row{m.Elem(i, jVal)},
			)
		} else {
			panic("Index operation not supported")
		}
	}

	return (&Matrix{}).Insert(
		Row{0},
	)
}

func (m *Matrix) s_i_index(i string, j int) *Matrix {
	switch len(i) {
	case 1:
		if i == ":" {
			returnMatrix := &Matrix{}
			for i := 0; i < m.horizontalDim; i += 1 {
				returnMatrix.Insert(
					Row{m.Table[i][j]},
				)
			}
			return returnMatrix
		}
		if iVal, err := strconv.Atoi(i); err == nil {
			return (&Matrix{}).Insert(
				Row{m.Elem(iVal, j)},
			)
		} else {
			panic("Index operation not supported")
		}
	}

	return (&Matrix{}).Insert(
		Row{0},
	)
}

func (m *Matrix) i_i_index(i int, j int) *Matrix {
	return (&Matrix{}).Insert(
		Row{m.matrixIndex(i, j)},
	)
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

	return (&Matrix{}).Insert(
		Row{0},
	)
}

//func (m *Matrix) rowVector(pos int) *Matrix {
//
//}

func (m *Matrix) colVector(pos int) *Matrix {
	if pos == -1 {
		returnMatrix := &Matrix{}
		for i := 0; i < m.horizontalDim * m.verticalDim; i += 1 {
			returnMatrix.Insert(
				Row{m.vectorIndex(i)},
			)
		}
		return returnMatrix
	}

	return (&Matrix{}).Insert(
		Row{0},
	)
}