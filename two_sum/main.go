package main

import "fmt"

func main()  {
	fmt.Println(TwoSum([]int{1, 2, 3}, 4))
}

func TwoSum(numbers []int, target int) [2]int {
	for i, val := range numbers {
		for i_in, val_in := range numbers[i+1:] {
			if val + val_in == target {
				return [2]int{ i, i_in + i + 1 }
			}
		}
	}

	return [2]int{}
}
