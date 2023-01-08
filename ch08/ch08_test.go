package ch08_test

import (
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch08"
)

func TestCountCommonNumbers(t *testing.T) {
	t.Parallel()

	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 3, 5, 7, 8, 9}

	want := 3 // 2,3,5
	if got := ch08.CountCommonNumbers(a, b); got != want {
		t.Errorf("got=%d, want=%d\n", got, want)
	}
}
