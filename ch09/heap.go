package ch09

import (
	"errors"
)

// Binary heap
type Heap[T any] struct {
	arr  []T
	less func(x, y T) bool
}

func NewHeap[T any](less func(x, y T) bool) *Heap[T] {
	return &Heap[T]{less: less}
}

func NewIntHeap() *Heap[int] {
	return NewHeap(func(x, y int) bool { return x < y })
}

func (h *Heap[T]) Push(nums ...T) {
	for _, num := range nums {
		h.push(num)
	}
}

func (h *Heap[T]) push(x T) {
	h.arr = append(h.arr, x)

	k := len(h.arr) - 1 // last index
	for k != 0 {
		p := (k - 1) / 2                // parent index
		if h.less(h.arr[p], h.arr[k]) { // child > parent
			h.arr[k], h.arr[p] = h.arr[p], h.arr[k] // swap parent for child
			k = p
		} else {
			break
		}
	}
}

// returns and removes top value
func (h *Heap[T]) Pop() (T, error) {
	top, err := h.Top()
	if err != nil {
		return *new(T), err
	}

	// remove(overwrite) top
	imax := len(h.arr) - 1
	h.arr[0] = h.arr[imax]
	h.arr = h.arr[:imax] // remove last element

	child1 := func(k int) int { return 2*k + 1 }
	child2 := func(k int) int { return 2*k + 2 }

	k := 0
	// c: child to swap for k(parent)
	for c := child1(k); c < len(h.arr); c = child1(k) { // while child1 exists
		if c2 := child2(k); c2 < len(h.arr) && h.less(h.arr[c], h.arr[c2]) { // if c2 exists & c1 value < c2 value
			c = c2
		}

		if h.less(h.arr[k], h.arr[c]) { // parent value < child value
			h.arr[k], h.arr[c] = h.arr[c], h.arr[k] // swap parent for child
			k = c
		} else {
			break
		}
	}
	return top, nil
}

func (h *Heap[T]) Top() (T, error) {
	if len(h.arr) == 0 {
		return *new(T), errors.New("empty heap")
	}
	return h.arr[0], nil
}

func (h *Heap[T]) Array() []T {
	return h.arr
}
