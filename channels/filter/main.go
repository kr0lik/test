package main

import "fmt"

func filter[T any](inputCh <-chan T, filterF func(T) bool) <-chan T {
	outputCh := make(chan T)

	go func() {
		defer close(outputCh)

		for v := range inputCh {
			if filterF(v) {
				outputCh <- v
			}
		}
	}()

	return outputCh
}

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}

		close(ch)
	}()

	filterF := func(x int) bool {
		if x%2 == 0 {
			return true
		}

		return false
	}

	resCh := filter(ch, filterF)

	for v := range resCh {
		fmt.Println(v)
	}
}
