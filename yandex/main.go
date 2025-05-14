package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.twitter.com",
		"https://www.youtube.com",
		"https://www.reddit.com",
		"https://www.facebook.com",
		"https://www.facebook.com",
	}

	for _, result := range requestLimiter(urls, 3) {
		fmt.Println(result)
	}
}

func requestLimiter(urls []string, limit int) []string {
	semaphore := NewSemaphore(limit)
	cache := NewCache(len(urls))

	wg := sync.WaitGroup{}

	for i, url := range urls {
		wg.Add(1)

		go func() {
			defer wg.Done()

			if cache.AddUrlAndCheckIsRequestedBefore(url) {
				return
			}

			semaphore.Lock()

			result := request(url, i)
			cache.AddUrlResult(url, result)

			semaphore.Unlock()
		}()
	}

	wg.Wait()

	res := make([]string, len(urls))

	for _, url := range urls {
		res = append(res, cache.GetResult(url))
	}

	return res
}

func request(url string, i int) string {
	time.Sleep(time.Second * 2)
	fmt.Println("request", url, i)
	return fmt.Sprintf("#%d: %s", i, url)
}

// Semaphore --------------------------------------------------------

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

// Cache --------------------------------------------------------

type Cache struct {
	mu       sync.RWMutex
	requests []string
	results  map[string]string
}

func NewCache(limit int) *Cache {
	return &Cache{requests: make([]string, limit), results: make(map[string]string)}
}

func (c *Cache) AddUrlAndCheckIsRequestedBefore(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.requests = append(c.requests, url)

	_, exist := c.results[url]
	if exist {
		return true
	}

	c.results[url] = ""
	return false
}

func (c *Cache) AddUrlResult(url string, result string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.results[url] = result
}

func (c *Cache) ListUrls() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.requests
}

func (c *Cache) GetResult(url string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.results[url]
}
