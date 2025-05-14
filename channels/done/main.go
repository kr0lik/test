package main

import (
	"fmt"
	"time"
)

func process(closeCh <-chan struct{}) <-chan struct{} {
	closeChDone := make(chan struct{})

	go func() {
		defer close(closeChDone)

		for {
			select {
			case <-closeCh:
				return
			default:
				fmt.Println("Processing")
			}
		}
	}()

	return closeChDone
}

func main() {
	closeCh := make(chan struct{})
	closeChDone := process(closeCh)

	time.Sleep(1 * time.Nanosecond)
	close(closeCh)
	<-closeChDone

	fmt.Println("Done")
}
