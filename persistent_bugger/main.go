package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Persistence(999))
}

func Persistence(n int) int {
	var counter int

	for {
		if n / 10 == 0 {
			break
		}
		helper(&n)
		counter++
	}

	return counter
}

func helper(n *int) {
	numbers := strconv.Itoa(*n)
	multiply := 1

	for index := range numbers {
		digit, _ := strconv.Atoi(string(numbers[index]))
		multiply *= digit
	}
	*n = multiply
}
