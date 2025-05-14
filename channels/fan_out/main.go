package main

import (
	"fmt"
	"sync"
)

func splitChannel[T any](inputCh <-chan T, n int) []<-chan T {
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

		count := 0
		for v := range inputCh {
			for {
				count++
				idx := count % n

				select {
				case outputChList[idx] <- v:
				default:
					continue
				}

				break
			}
		}
	}()

	outputCastChList := make([]<-chan T, n)
	for i, outputCh := range outputChList {
		outputCastChList[i] = outputCh
	}

	return outputCastChList
}

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	channels := splitChannel(ch, 2)

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
