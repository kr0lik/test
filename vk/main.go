package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var ErrQueueEmpty = errors.New("queue empty")

type Notification struct {
	text string
	time time.Time
}

func NewNotification(text string, time time.Time) Notification {
	return Notification{
		text: text,
		time: time,
	}
}

func (n *Notification) Text() string {
	return n.text
}

func (n *Notification) Time() time.Time {
	return n.time
}

type NotificationQueue struct {
	queue []*Notification
	mu    sync.Mutex
}

func NewNotificationQueue() *NotificationQueue {
	return &NotificationQueue{}
}

func (nq *NotificationQueue) AddNotification(n Notification) {
	nq.mu.Lock()
	nq.queue = append(nq.queue, &n)
	nq.mu.Unlock()
}

func (nq *NotificationQueue) PopNotification() (Notification, error) {
	nq.mu.Lock()
	defer nq.mu.Unlock()

	if len(nq.queue) == 0 {
		return Notification{}, ErrQueueEmpty
	}

	res := nq.queue[0]

	nq.queue = nq.queue[1:]

	return *res, nil
}

func (nq *NotificationQueue) GetRandomNotification() (Notification, error) {
	nq.mu.Lock()
	defer nq.mu.Unlock()

	if len(nq.queue) == 0 {
		return Notification{}, ErrQueueEmpty
	}

	index := rand.Intn(len(nq.queue))
	return *nq.queue[index], nil
}

func main() {
	fmt.Println("TestAddNotification")
	TestAddNotification()

	fmt.Println("TestPopNotification")
	TestPopNotification()

	fmt.Println("TestGetRandomNotification")
	TestGetRandomNotification()
}

func TestAddNotification() {
	const op = "TestAddNotification"

	q, testData := getQueue()

	if len(q.queue) != len(testData) {
		log.Fatalf("%s: expected queue length %d, got %d", op, len(testData), len(q.queue))
	}

	for i, n := range testData {
		storedNotification := q.queue[i]

		if storedNotification.Text() != n.Text() || storedNotification.Time() != n.Time() {
			log.Fatalf("%s: notification not added correctly", op)
		}
	}
}

func TestPopNotification() {
	const op = "TestPopNotification"

	q, testData := getQueue()

	for i, n := range testData {
		pop, err := q.PopNotification()
		if err != nil {
			log.Fatalf("%s: %s", op, err.Error())
		}

		if pop.Text() != n.Text() {
			log.Fatalf("%s:  Text %s, got %s", op, n.Text(), pop.Text())
		}

		expectedLen := len(testData) - i - 1
		if len(q.queue) != expectedLen {
			log.Fatalf("%s: expected queue length %d, got %d", op, expectedLen, len(q.queue))
		}
	}
}

func TestGetRandomNotification() {
	const op = "TestGetRandomNotification"

	q, testData := getQueue()

	for i := 0; i < len(q.queue); i++ {
		_, err := q.GetRandomNotification()
		if err != nil {
			log.Fatalf("%s: %s", op, err.Error())
		}
	}

	if len(q.queue) != len(testData) {
		log.Fatalf("%s: expected queue length %d, got %d", op, len(testData), len(q.queue))
	}
}

func getQueue() (*NotificationQueue, []Notification) {
	testData := []Notification{
		NewNotification("First notification", time.Now()),
		NewNotification("Second notification", time.Now()),
		NewNotification("Third notification", time.Now()),
	}

	q := NewNotificationQueue()

	for _, n := range testData {
		q.AddNotification(n)
	}

	return q, testData
}
