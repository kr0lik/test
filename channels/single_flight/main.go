package main

import (
	"fmt"
	"sync"
	"time"
)

type call struct {
	value interface{}
	err   error
	done  chan struct{}
}

type SingleFlight struct {
	mu   sync.Mutex
	list map[string]*call
}

func NewSingleFlight() *SingleFlight {
	return &SingleFlight{list: make(map[string]*call)}
}

func (sf *SingleFlight) Do(key string, f func() (interface{}, error)) (interface{}, error) {
	sf.mu.Lock()
	cal, exist := sf.list[key]
	sf.mu.Unlock()

	if exist {
		<-cal.done
		return cal.value, cal.err
	}

	cal = &call{done: make(chan struct{})}

	sf.mu.Lock()
	sf.list[key] = cal
	sf.mu.Unlock()

	go func() {
		value, err := f()

		cal.value = value
		cal.err = err
		close(cal.done)

		sf.mu.Lock()
		delete(sf.list, key)
		sf.mu.Unlock()
	}()

	<-cal.done
	return cal.value, cal.err
}

func main() {
	wg := sync.WaitGroup{}

	singleFlight := NewSingleFlight()

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			value, err := singleFlight.Do("key", func() (interface{}, error) {
				time.Sleep(time.Second)
				fmt.Println("Do", i)
				return fmt.Sprintf("Return %d", i), nil
			})

			fmt.Println(i, ":", value, err)
		}()
	}

	wg.Wait()
}
