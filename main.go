package main

import (
	"fmt"
	. "github.com/JonathanRosado/Shako/linalg/matrix"
	//. "github.com/JonathanRosado/Shako/ml"
)

func main() {
	fmt.Println("hello world")


	m := Matrix{}

	m.Insert(
		Row{1,2,3,4},
		Row{12,32,43,2},
		Row{6,3,45,53},
	)

	m.Index(0, ":", func(i int, j int, elem float64) float64 {
		return elem*elem
	}).Print()

}