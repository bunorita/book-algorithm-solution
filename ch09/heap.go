package ch09

import (
	"errors"
)

// Binary heap
type Heap []int

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) Push(nums ...int) {
	for _, num := range nums {
		h.push(num)
	}
}

func (h *Heap) push(x int) {
	*h = append(*h, x)

	k := len(*h) - 1 // last index
	for k != 0 {
		p := (k - 1) / 2       // parent index
		if (*h)[k] > (*h)[p] { // child > parent
			(*h)[k], (*h)[p] = (*h)[p], (*h)[k] // swap parent for child
			k = p
		} else {
			break
		}
	}
}

// returns and removes top value
func (h *Heap) Pop() (int, error) {
	top, err := h.Top()
	if err != nil {
		return 0, err
	}

	// remove(overwrite) top
	imax := len(*h) - 1
	(*h)[0] = (*h)[imax]
	(*h) = (*h)[:imax] // remove last element

	// h = [1,8,3,0,4]

	child1 := func(k int) int { return 2*k + 1 }
	child2 := func(k int) int { return 2*k + 2 }

	k := 0
	// c: child to swap for k(parent)
	for c := child1(k); c < len(*h); c = child1(k) { // while child1 exists
		if c2 := child2(k); c2 < len(*h) && (*h)[c] < (*h)[c2] { // if c2 exists & c1 value < c2 value
			c = c2
		}

		if (*h)[k] < (*h)[c] { // parent value < child value
			(*h)[k], (*h)[c] = (*h)[c], (*h)[k] // swap parent for child
			k = c
		} else {
			break
		}
	}
	return top, nil
}

func (h *Heap) Top() (int, error) {
	if len(*h) == 0 {
		return 0, errors.New("empty heap")
	}
	return (*h)[0], nil
}

func (h *Heap) Array() []int {
	return *h
}
