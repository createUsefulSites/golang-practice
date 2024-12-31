package main

import "fmt"

func main()  {
	arr := []int{124, 0, -1, 12}

	fmt.Println(sortArrays(arr))
}

func sortArrays(arr []int) [2]int {
	for x := 0; x < len(arr) - 1; x++ {
		for y := 0; y < len(arr) - 1 - x; y++ {
			if arr[y] > arr[y + 1] {
				arr[y], arr[y + 1] = arr[y + 1], arr[y]
			}
		}
	}

	return [2]int(arr[len(arr)-2:])
}
