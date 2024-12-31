package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(alphanumeric("asf"))
}

func alphanumeric(str string) bool {
  matched, err := regexp.MatchString(`^[a-zA-Z0-9]+$`, str)
	if err != nil {
		return false
	}

	return matched
}
