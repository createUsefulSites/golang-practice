package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Is_valid_ip(""))
	fmt.Println(Is_valid_ip("123.45.67.89"))
}

func Is_valid_ip(ip string) bool {
	slice := strings.Split(ip, ".")
	if len(slice) != 4 {
		return false
	}

	for _, part := range slice {
		if part[0] == '0' && len(part) > 1 {
			return false
		}

		parsedInt, err := strconv.Atoi(part)
		if err != nil {
			return false
		}

		if parsedInt < 0 || parsedInt > 255 {
			return false
		}
	}

	return true
}
