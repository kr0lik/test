package main

import (
	"fmt"
	"time"
)

func worker[T any](jobs <-chan T, results chan<- T) {
	for j := range jobs {
		time.Sleep(time.Second)
		results <- j
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 1; i <= 3; i++ {
		go worker(jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for j := 1; j <= 5; j++ {
		res := <-results
		fmt.Println(res)
	}
	close(results)
}
