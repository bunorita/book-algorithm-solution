package ch07_test

import (
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch07"
)

func TestIntervalScheduling(t *testing.T) {
	t.Parallel()

	// N=3
	intervals := []*ch07.Interval{
		{1, 5},
		{4, 6},
		{6, 8},
	}
	want := 2
	if got := ch07.IntervalScheduling(intervals); got != want {
		t.Errorf("got=%d, want=%d\n", got, want)
	}

}
