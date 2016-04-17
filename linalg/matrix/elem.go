package matrix

import (
	//log "github.com/Sirupsen/logrus"
)

func (m *Matrix) Elem(args ...int) float64 {

	switch len(args) {
	case 0:
		return m.Elem(0)
	case 1:
		return m.vectorIndex(args[0])
	case 2:
		return m.matrixIndex(args[0], args[1])
	default:
		panic("Too many arguments given to Matrix.Elem")
	}

	return -1
}