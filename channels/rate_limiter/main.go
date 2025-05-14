package main

import (
	"fmt"
	"time"
)

type ReteLimiter struct {
	leakyBucketCh chan struct{}
	stopCh        chan struct{}
}

func NewReteLimiter(limit int, period time.Duration) *ReteLimiter {
	rl := &ReteLimiter{leakyBucketCh: make(chan struct{}), stopCh: make(chan struct{})}
	interval := period.Nanoseconds() / int64(limit)

	go rl.leak(time.Duration(interval))

	return rl
}

func (rl *ReteLimiter) Allow() bool {
	select {
	case rl.leakyBucketCh <- struct{}{}:
		return true
	default:
		return false
	}
}

func (rl *ReteLimiter) Stop() {
	close(rl.stopCh)
}

func (rl *ReteLimiter) leak(interval time.Duration) {
	<-rl.leakyBucketCh

	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	defer close(rl.leakyBucketCh)

	for {
		select {
		case <-ticker.C:
			<-rl.leakyBucketCh
		case <-rl.stopCh:
			return
		}
	}
}

func main() {
	rateLimiter := NewReteLimiter(1, time.Second)

	for i := 0; i < 20; i++ {
		time.Sleep(time.Second / 2)
		if rateLimiter.Allow() {
			fmt.Println(i, "is allowed")
		} else {
			fmt.Println(i, "is NOT allowed")
		}
	}

	rateLimiter.Stop()
}
