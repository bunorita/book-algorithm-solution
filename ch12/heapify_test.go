package ch12

import (
	"reflect"
	"testing"
)

func TestHeapify(t *testing.T) {
	t.Parallel()

	a := []int{1, 2, 3, 4}
	want := []int{4, 2, 3, 1} // heap

	n := len(a)
	for i := (n - 1) / 2; i >= 0; i-- {
		heapify(&a, i, n)
	}

	if got := a; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
