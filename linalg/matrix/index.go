package matrix

import (
	//"reflect"
)
import "github.com/JonathanRosado/Shako/linalg/matrix/data"

func (m *Matrix) Index(args ...interface{}) *Matrix {

	switch len(args) {
	case 1:
		switch arg0Val := args[0].(type) {
		case string:
			if arg0Val == ":" {
				return m.columnVector()
			} else {
				panic("Argument not supported by Matrix.Index")
			}
		case int:
			matrix := &Matrix{}
			matrix.data = *(&data.Data{}).Create([][]float64{
				[]float64{m.vectorIndex(arg0Val)},
			})
			return matrix
		default:
			panic("Argument not supported by Matrix.Index")
		}
	case 2:
		switch arg0Val := args[0].(type) {
		case string:
			switch arg1Val := args[1].(type) {
			case string:
				return m.s_s_index(arg0Val, arg1Val)
			case int:
				return m.s_i_index(arg0Val, arg1Val)
			default:
				panic("Argument not supported by Matrix.Index")
			}
		case int:
			switch arg1Val := args[1].(type) {
			case string:
				return m.i_s_index(arg0Val, arg1Val)
			case int:
				return m.i_i_index(arg0Val, arg1Val)
			default:
				panic("Argument not supported by Matrix.Index")
			}
		default:
			panic("Argument not supported by Matrix.Index")
		}
	case 3:
		returnMatrix := &Matrix{}

		switch arg0Val := args[0].(type) {
		case string:
			switch arg1Val := args[1].(type) {
			case string:
				returnMatrix = m.s_s_index(arg0Val, arg1Val)
			case int:
				returnMatrix = m.s_i_index(arg0Val, arg1Val)
			default:
				panic("Argument not supported by Matrix.Index")
			}
		case int:
			switch arg1Val := args[1].(type) {
			case string:
				returnMatrix = m.i_s_index(arg0Val, arg1Val)
			case int:
				returnMatrix = m.i_i_index(arg0Val, arg1Val)
			default:
				panic("Argument not supported by Matrix.Index")
			}
		default:
			panic("Argument not supported by Matrix.Index")
		}

		returnMatrix.ElemOp(args[2].(func(int, int, float64) (float64)))

		return returnMatrix
	default:
		panic("Too many arguments given to Matrix.Index")
	}

	matrix := &Matrix{}
	matrix.data = *(&data.Data{}).Create([][]float64{
		[]float64{0},
	})
	return matrix
}
