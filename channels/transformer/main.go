package main

import "fmt"

func transform[T any](inputCh <-chan T, transformF func(T) T) <-chan T {
	outputCh := make(chan T)

	go func() {
		defer close(outputCh)

		for v := range inputCh {
			outputCh <- transformF(v)
		}
	}()

	return outputCh
}

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	transformFunc := func(x int) int {
		return x * 10
	}

	resCh := transform(ch, transformFunc)

	for v := range resCh {
		fmt.Println(v)
	}
}
