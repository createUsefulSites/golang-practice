package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FirstNonRepeating("sTreSS"))
}

func FirstNonRepeating(str string) string {
	lettersMap := make(map[string]int)
	originalCaseMap := make(map[string]string)
	keysOrder := []string{}

	for _, val := range strings.Split(str, "") {
		lowerKey := strings.ToLower(val)
		if _, ok := originalCaseMap[lowerKey]; !ok {
			originalCaseMap[lowerKey] = val
		}

		if _, ok := lettersMap[lowerKey]; !ok {
			lettersMap[lowerKey] = 1
			keysOrder = append(keysOrder, lowerKey)
		} else {
			lettersMap[lowerKey] += 1
		}
	}

	for _, key := range keysOrder {
		if lettersMap[key] == 1 {
			return originalCaseMap[key]
		}
	}

	return ""
}