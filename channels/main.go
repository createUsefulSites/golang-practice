package main

import (
	"fmt"
	"time"
)

func main() {
	chanel := make(chan int, 2)
	
	go func ()  {
		time.Sleep(time.Second)
		chanel <- 123
		chanel <- 234
	}()

	fmt.Println(<-chanel)
	chanel <- 2341242
	fmt.Println(<-chanel)
	fmt.Println(<-chanel)
}
