package ch08

import (
	"errors"
	"fmt"
	"strconv"
)

type IQueue interface {
	Enqueue(int) error
	Dequeue() (int, error)
	Values() []int
}

// const maxQueueSize = 100000

const maxQueueSize = 5

var _ IQueue = (*Queue)(nil)

type Queue struct {
	head, tail int
	values     [maxQueueSize]int // ring(circular) buffer
}

func NewQueue() *Queue {
	return &Queue{}
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

func (q *Queue) IsEmpty() bool {
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
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	x := q.values[q.head]
	q.head++
	if q.head == maxQueueSize {
		q.head = 0
	}
	return x, nil
}

var _ IQueue = (*LinkedListQueue)(nil)

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{values: NewLinkedList()}
}

type LinkedListQueue struct {
	values *LinkedList
}

func (q *LinkedListQueue) Values() []int {
	size := q.values.Size()
	values := make([]int, size)
	for i := 0; i < size; i++ {
		x, _ := strconv.Atoi(q.values.GetNode(i).value)
		values[i] = x
	}
	return values
}

func (q *LinkedListQueue) Enqueue(x int) error {
	q.values.Append(NewLinkedListNode("", strconv.Itoa(x)))
	return nil
}

func (q *LinkedListQueue) Dequeue() (int, error) {
	first := q.values.GetNode(0)
	q.values.Erase(first)
	x, _ := strconv.Atoi(first.value)
	return x, nil
}
