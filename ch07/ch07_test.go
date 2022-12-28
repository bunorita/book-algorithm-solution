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

func TestMultipleArray(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name string
		a    []int
		b    []int
		want int
	}{
		{
			name: "N=3",
			a:    []int{3, 2, 9},
			b:    []int{5, 7, 4},
			want: 7,
		},
		{
			name: "N=7",
			a:    []int{3, 4, 5, 2, 5, 5, 9},
			b:    []int{1, 1, 9, 6, 3, 8, 7},
			want: 22,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := ch07.MultipleArray(tt.a, tt.b); got != tt.want {
				t.Errorf("got=%d, want=%d\n", got, tt.want)
			}
		})
	}
}
