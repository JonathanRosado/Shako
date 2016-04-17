package matrix

import (
	"fmt"
	"reflect"
)

type Row []float64

type ElemWiseFunc func(int, int, float64) (float64)

type Matrix struct {
	Table  []Row
	cols int
	rows int
}

func (m *Matrix) Print() {
	for index := range m.Table {
		fmt.Printf("%f\n", m.Table[index])
	}
}

func (m *Matrix) ElemOp(fn ElemWiseFunc) *Matrix {
	for i := 0; i < m.rows; i += 1 {
		for j := 0; j < m.cols; j += 1 {
			m.Table[i][j] = fn(i, j, m.Table[i][j])
		}
	}
	return m
}

func (m *Matrix) ElemOpInterface(fn interface{}) *Matrix {
	for i := 0; i < m.rows; i += 1 {
		for j := 0; j < m.cols; j += 1 {
			value := reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(i), reflect.ValueOf(j), reflect.ValueOf(m.Table[i][j])})[0]

			m.Table[i][j] = value.Float()
		}
	}
	return m
}

