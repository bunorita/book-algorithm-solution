package ch08

import (
	"errors"
	"fmt"
	"strconv"
)

type IQueue[T any] interface {
	Enqueue(T) error
	Dequeue() (T, error)
	Values() []T
}

// const maxQueueSize = 100000

const maxQueueSize = 10

var _ IQueue[int] = (*Queue[int])(nil)

type Queue[T any] struct {
	head, tail int
	values     [maxQueueSize]T // ring(circular) buffer
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Inspect() {
	fmt.Printf("head=%d, tail=%d, values=%v\n", q.head, q.tail, q.values)
}

func (q *Queue[T]) Values() []T {
	if q.head < q.tail {
		return q.values[q.head:q.tail]
	} else {
		front := make([]T, maxQueueSize-q.head)
		copy(front, q.values[q.head:]) // Not to overwrite q.values
		return append(front, q.values[:q.tail]...)
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head == q.tail
}

func (q *Queue[T]) isFull() bool {
	return (q.tail+1)%maxQueueSize == q.head
}

func (q *Queue[T]) Enqueue(x T) error {
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

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return *new(T), errors.New("queue is empty")
	}
	x := q.values[q.head]
	q.head++
	if q.head == maxQueueSize {
		q.head = 0
	}
	return x, nil
}

func (q *Queue[T]) Push(x T) error {
	return q.Enqueue(x)
}

func (q *Queue[T]) Pop() (T, error) {
	return q.Dequeue()
}

var _ IQueue[int] = (*LinkedListQueue)(nil)

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
