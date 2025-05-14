package main

import (
	"fmt"
	"sync"
	"time"
)

type ErrGroup struct {
	err  error
	wg   sync.WaitGroup
	once sync.Once

	doneCh chan struct{}
}

func NewErrGroup() (*ErrGroup, chan struct{}) {
	doneCh := make(chan struct{})
	return &ErrGroup{doneCh: doneCh}, doneCh
}

func (e *ErrGroup) Do(f func() error) {
	e.wg.Add(1)

	go func() {
		defer e.wg.Done()

		select {
		case <-e.doneCh:
			return
		default:
			if err := f(); err != nil {
				e.once.Do(func() {
					e.err = err
					close(e.doneCh)
				})
			}
		}
	}()
}

func (e *ErrGroup) Wait() error {
	e.wg.Wait()
	return e.err
}

func main() {
	wg, wgDoneCh := NewErrGroup()
	for i := 0; i < 10; i++ {
		wg.Do(func() error {
			time.Sleep(time.Second)

			select {
			case <-wgDoneCh:
				fmt.Println("cancel")
				return fmt.Errorf("done")
			default:
				return fmt.Errorf("fail")
			}
		})
	}

	if err := wg.Wait(); err != nil {
		fmt.Println(err)
	}
}
