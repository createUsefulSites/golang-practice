package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	numbersChannel := make(chan int)
	var wg sync.WaitGroup

	timeStart := time.Now().UnixMilli()

	wg.Add(1)
	go generateRandomNumber(numbersChannel, &wg)

	wg.Add(1)
	go func() {
		for num := range numbersChannel {
			fmt.Println(num)
		}

		wg.Done()
	}()

	wg.Wait()

	timeEnd := time.Now().UnixMilli()
	fmt.Println("Время выполнения:", (timeEnd - timeStart), "мс")
}

func generateRandomNumber(write chan<- int, wg *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000000; i++ {
		write <- rand.Intn(1000)
	}

	close(write)
	wg.Done()
}