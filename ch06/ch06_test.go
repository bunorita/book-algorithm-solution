package ch06_test

import (
	"fmt"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch06"
)

func TestBinarySearch(t *testing.T) {
	t.Parallel()

	a := []int{3, 5, 8, 10, 14, 17, 21, 39}

	tests := []*struct {
		key  int
		want int
	}{
		{key: 10, want: 3},
		{key: 3, want: 0},
		{key: 39, want: 7},
		{key: 9, want: -1},
		{key: 100, want: -1},
		{key: -100, want: -1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("key=%d", tt.key), func(t *testing.T) {
			t.Parallel()

			if got := ch06.BinarySearch(a, tt.key); got != tt.want {
				t.Errorf("got=%d, want=%d\n", got, tt.want)
			}
		})
	}
}

func TestBinarySearchGen(t *testing.T) {
	t.Parallel()

	a := []int{3, 5, 8, 10, 14, 17, 21, 39}

	tests := []*struct {
		key  int
		want int
	}{
		{key: 10, want: 3},
		{key: 3, want: 0},
		{key: 39, want: 7},
		{key: 9, want: 3},
		{key: 100, want: 8}, // a[i] >= 100 となるiは存在しない
		{key: -100, want: 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("key=%d", tt.key), func(t *testing.T) {
			t.Parallel()

			if got := ch06.BinarySearchGen(a, tt.key); got != tt.want {
				t.Errorf("got=%d, want=%d\n", got, tt.want)
			}
		})
	}
}
