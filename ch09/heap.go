package ch09

import "errors"

// Binary heap
type Heap []int

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) Push(x int) {
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

	if len(*h) == 0 {
		*h = []int{}
	} else {
		// remove(overwrite) top
		imax := len(*h) - 1
		(*h)[0] = (*h)[imax]
		(*h) = (*h)[:imax] // remove last element

		imax = len(*h) - 1 // new value
		k := 0
		for k != imax {
			c := 2*k + 2           // right child
			if (*h)[k] < (*h)[c] { // parent < child
				(*h)[k], (*h)[c] = (*h)[c], (*h)[k] // swap parent for child
				k = c
			} else {
				break
			}
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
