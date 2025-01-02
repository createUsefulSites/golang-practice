package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		for {
			chan1 <- "сообщение записано, прошло 200мс"
			time.Sleep(time.Second / 5)
		}
	}()

	go func() {
		for {
			chan2 <- "сообщение записано, прошло 1000мс"
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case message := <-chan1:
			fmt.Println("info from channel 1: ", message)
		case message := <-chan2:
			fmt.Println("info from channel 2: ", message)
		}
	}
}
