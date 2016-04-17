package matrix

import "testing"

func BenchmarkIndex1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m.Index(3)
	}
}

func BenchmarkIndex2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m.Index(2,3)
	}
}

func BenchmarkIndex3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m.Index(2,"3")
	}
}

func BenchmarkIndex4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m.Index("2",3)
	}
}

func BenchmarkIndex5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m.Index(":",3)
	}
}

func BenchmarkIndex6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m.Index(2,":")
	}
}

func BenchmarkIndex7(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m.Index(":")
	}
}

func BenchmarkIndex8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		m.Index(0, ":", func(i int, j int, elem float64) float64 {
			return elem*elem
		})
	}
}

