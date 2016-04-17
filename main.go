package main

import (
	"fmt"
	. "github.com/JonathanRosado/Shako/linalg/matrix"
	//"github.com/JonathanRosado/Shako/linalg/matrix/data"
	//. "github.com/JonathanRosado/Shako/ml"
)

func main() {
	fmt.Println("hello world")


	m := (&Matrix{}).Create(
		[]float64{1,2,3,4},
		[]float64{12,32,43,2},
		[]float64{6,3,45,53},
	)

	m.Index(0, ":", func(i int, j int, elem float64) float64 {
		return elem*elem
	}).Print()

}