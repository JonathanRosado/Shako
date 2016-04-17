package matrix

import (
	"github.com/JonathanRosado/Shako/linalg/matrix/data"
)

type ElemWiseFunc func(int, int, float64) (float64)


type Matrix struct {
	data data.Data
}

func (m *Matrix) Print() {
	m.data.Print()
}

func (m *Matrix) ElemOp(fn ElemWiseFunc) *Matrix {
	for i := 0; i < m.data.Rows; i += 1 {
		for j := 0; j < m.data.Cols; j += 1 {
			m.data.Set(i, j, fn(i, j, m.data.Get(i, j)))
		}
	}
	return m
}

//func (m *Matrix) ElemOpInterface(fn interface{}) *Matrix {
//	for i := 0; i < m.rows; i += 1 {
//		for j := 0; j < m.cols; j += 1 {
//			value := reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(i), reflect.ValueOf(j), reflect.ValueOf(m.Table[i][j])})[0]
//
//			m.Table[i][j] = value.Float()
//		}
//	}
//	return m
//}

