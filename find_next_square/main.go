package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindNextSquare(12))
	fmt.Println(FindNextSquare(144))
}

func FindNextSquare(value int64) int64 {
	square := math.Sqrt(float64(value))
	if math.Pow(math.Floor(square), 2) == math.Pow(square, 2) {
		return int64(math.Pow(square + 1, 2))
	}

  return -1
}
