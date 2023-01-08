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

func TestCountCommonNumberPairs(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name string
		a, b []int
		want int
	}{
		{
			name: "case1",
			a:    []int{1},
			b:    []int{2},
			want: 0,
		},
		{
			name: "case2",
			a:    []int{1},
			b:    []int{2, 1},
			want: 1, // (i,j) = (0,1)
		},
		{
			name: "case3",
			a:    []int{1, 2, 3},
			b:    []int{2, 2, 1, 4},
			want: 3, // (i,j) = (0,2), (1,0), (1,1)
		},
		{
			name: "case4", // swapped a, b of case3
			a:    []int{2, 2, 1, 4},
			b:    []int{1, 2, 3},
			want: 3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := ch08.CountCommonNumberPairs(tt.a, tt.b); got != tt.want {
				t.Errorf("got=%d, want=%d\n", got, tt.want)
			}
		})
	}
}
