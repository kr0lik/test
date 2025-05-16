package main

import (
	"fmt"
	"time"
)

func hello(n int, ch chan struct{}) {
	for i := 0; i < n; i++ {
		fmt.Print("Hallow")
		ch <- struct{}{}
		<-ch
	}
}

func world(n int, ch chan struct{}) {
	for i := 0; i < n; i++ {
		<-ch
		fmt.Println("World")
		ch <- struct{}{}
	}
}

func main() {
	n := 10

	ch := make(chan struct{})

	go hello(n, ch)
	go world(n, ch)

	time.Sleep(time.Second)
}
