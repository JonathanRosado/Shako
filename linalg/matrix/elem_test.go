package matrix

import "testing"

var m *Matrix

func init() {
	m = (&Matrix{}).Create(
		[]float64{1,2,3,4},
		[]float64{12,32,43,2},
		[]float64{6,3,45,53},
	)

}

func BenchmarkElem1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m.Elem(3)
	}
}

func BenchmarkElem2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m.Elem(2,3)
	}
}
