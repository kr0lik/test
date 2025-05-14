package main

import (
	"fmt"
	"sync"
)

type Barrier struct {
	mu    sync.Mutex
	size  int
	count int

	beforeCh chan struct{}
	afterCh  chan struct{}
}

func NewBarrier(size int) *Barrier {
	return &Barrier{size: size, beforeCh: make(chan struct{}, size), afterCh: make(chan struct{}, size)}
}

func (b *Barrier) Before() {
	b.mu.Lock()
	b.count++

	if b.count == b.size {
		for i := 0; i < b.size; i++ {
			b.beforeCh <- struct{}{}
		}
	}

	b.mu.Unlock()

	<-b.beforeCh
}

func (b *Barrier) After() {
	b.mu.Lock()
	b.count--

	if b.count == 0 {
		for i := 0; i < b.size; i++ {
			b.afterCh <- struct{}{}
		}
	}

	b.mu.Unlock()

	<-b.afterCh
}

func main() {
	wg := sync.WaitGroup{}

	count := 3
	barrier := NewBarrier(count)
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < count; j++ {
				barrier.Before()
				start()

				barrier.After()
				stop()
			}
		}()
	}

	wg.Wait()
}

func start() {
	fmt.Println("start")
}

func stop() {
	fmt.Println("stop")
}
