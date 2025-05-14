package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	buffer chan struct{}
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{make(chan struct{}, n)}
}

func (s *Semaphore) Lock() {
	s.buffer <- struct{}{}
}

func (s *Semaphore) Unlock() {
	<-s.buffer
}

func main() {
	wg := &sync.WaitGroup{}

	semaphore := NewSemaphore(2)

	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()

			semaphore.Lock()
			time.Sleep(time.Second * 2)
			fmt.Println(i)
			semaphore.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Done")
}
