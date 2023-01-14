package ch08

import (
	"errors"
	"fmt"
)

const maxQueueSize = 100000

// const maxQueueSize = 5

type Queue struct {
	head, tail int
	values     [maxQueueSize]int // ring(circular) buffer
}

func NewQueue(nums ...int) (*Queue, error) {
	q := &Queue{}
	for _, num := range nums {
		if err := q.Enqueue(num); err != nil {
			return nil, err
		}
	}
	return q, nil
}

func (q *Queue) Inspect() {
	fmt.Printf("head=%d, tail=%d, values=%v\n", q.head, q.tail, q.values)
}

func (q *Queue) Values() []int {
	if q.head < q.tail {
		return q.values[q.head:q.tail]
	} else {
		front := make([]int, maxQueueSize-q.head)
		copy(front, q.values[q.head:]) // Not to overwrite q.values
		return append(front, q.values[:q.tail]...)
	}
}

func (q *Queue) isEmpty() bool {
	return q.head == q.tail
}

func (q *Queue) isFull() bool {
	return (q.tail+1)%maxQueueSize == q.head
}

func (q *Queue) Enqueue(x int) error {
	if q.isFull() {
		return errors.New("queue is full")
	}
	q.values[q.tail] = x
	q.tail++
	if q.tail == maxQueueSize {
		q.tail = 0
	}
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("queue is empty")
	}
	x := q.values[q.head]
	q.head++
	if q.head == maxQueueSize {
		q.head = 0
	}
	return x, nil
}