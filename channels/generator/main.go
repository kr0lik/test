package main

import (
	"fmt"
	"time"
)

func Generate(start int, end int) <-chan int {
	outputCh := make(chan int)

	go func() {
		defer close(outputCh)

		for i := start; i <= end; i++ {
			time.Sleep(time.Second)
			outputCh <- i
		}
	}()

	return outputCh
}

func main() {
	for number := range Generate(1, 10) {
		fmt.Println(number)
	}
}
