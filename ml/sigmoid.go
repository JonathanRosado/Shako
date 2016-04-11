package ml

import "math"

func Sigmoid(z float64) float64 {
	return 1.0/(1.0 + math.Pow(math.E, -(z)))
}
