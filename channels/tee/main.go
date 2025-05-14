package main

import (
	"fmt"
	"sync"
)

func tee[T any](inputCh <-chan T, n int) []<-chan T {
	outputChList := make([]chan T, n)
	for i, _ := range outputChList {
		outputChList[i] = make(chan T)
	}

	go func() {
		defer func() {
			for _, ch := range outputChList {
				close(ch)
			}
		}()

		for v := range inputCh {
			for i := 0; i < n; i++ {
				outputChList[i] <- v
			}
		}
	}()

	outputChCastList := make([]<-chan T, n)
	for i, outputCh := range outputChList {
		outputChCastList[i] = outputCh
	}

	return outputChCastList
}

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	channels := tee(ch, 2)

	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	go func() {
		defer wg.Done()

		for v := range channels[0] {
			fmt.Println("ch0", v)
		}
	}()

	go func() {
		defer wg.Done()

		for v := range channels[1] {
			fmt.Println("ch1", v)
		}
	}()

	wg.Wait()
}
