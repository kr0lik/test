package main

import (
	"fmt"
	"sync"
)

func mergeChannels[T any](channels ...<-chan T) <-chan T {
	var outputCh = make(chan T)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, channel := range channels {
		go func() {
			defer wg.Done()

			for v := range channel {
				outputCh <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outputCh)
	}()

	return outputCh
}

func main() {
	chList := make([]chan string, 3)
	for i, _ := range chList {
		chList[i] = make(chan string)
	}

	go func() {
		defer func() {
			for _, ch := range chList {
				close(ch)
			}
		}()

		for i := 0; i < 10; i++ {
			for ci, ch := range chList {
				ch <- fmt.Sprintf("%d: %s-%d", i, "ch", ci)
			}
		}
	}()

	castChList := make([]<-chan string, len(chList))
	for i, outputCh := range chList {
		castChList[i] = outputCh
	}

	for v := range mergeChannels(castChList...) {
		fmt.Println(v)
	}
}
