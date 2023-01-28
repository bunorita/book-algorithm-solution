package ch09_test

import (
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch09"
)

func TestHeap(t *testing.T) {
	t.Parallel()

	h := ch09.NewHeap()
	for _, x := range []int{5, 3, 7, 1} {
		h.Push(x)
	}

	got, err := h.Top()
	if err != nil {
		t.Fatal(err)
	}
	if want := 7; got != want {
		t.Errorf("got=%d, want=%d", got, want)
	}

	got, err = h.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if want := 7; got != want {
		t.Errorf("got=%d, want=%d", got, want)
	}

	got, err = h.Top()
	if err != nil {
		t.Fatal(err)
	}
	if want := 5; got != want {
		t.Errorf("got=%d, want=%d", got, want)
	}

	h.Push(11)
	got, err = h.Top()
	if err != nil {
		t.Fatal(err)
	}
	if want := 11; got != want {
		t.Errorf("got=%d, want=%d", got, want)
	}
}
