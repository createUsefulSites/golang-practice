package main

import (
	"fmt"
)

func main()  {
	var arr = []int{160, 18, -19, 18, 22}

	fmt.Println(FindOutlier(arr))
}

func absInt(num int) int {
	if num < 0 {
			return -num
	}
	return num
}

func FindItemHelper(arr []int, equal int) int {
	for _, item := range arr {
		if absInt(item) % 2 == equal {
			return item
		}
	}
	return 0
}

func FindOutlier(integers []int) int {
	if isItemsOdd([3]int(integers[:3])) {
		return FindItemHelper(integers, 0)
	}
	return FindItemHelper(integers, 1)
}

func isItemsOdd(arr [3]int) bool {
	var sumValue int = 0
	for x := 0; x < len(arr); x++ {
		sumValue += absInt(arr[x]) % 2
	}

	return sumValue > 1
}
