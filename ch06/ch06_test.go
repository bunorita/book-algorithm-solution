package ch06_test

import (
	"fmt"
	"reflect"
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

func TestMinSum(t *testing.T) {
	t.Parallel()

	a := []int{4, 3, 12, 7, 11}
	b := []int{8, 1, 5, 2, 9}

	tests := []*struct {
		k    int
		want int
	}{
		{k: 0, want: 4},
		{k: 4, want: 4},
		{k: 5, want: 5},
		{k: 6, want: 6},
		{k: 7, want: 8},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("k=%d", tt.k), func(t *testing.T) {
			t.Parallel()

			if got := ch06.MinSum(a, b, tt.k); got != tt.want {
				t.Errorf("want %d, but got %d\n", tt.want, got)
			}
		})
	}
}

func TestShooting(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name string
		h    []int
		s    []int
		want int
	}{
		{
			name: "N=4",
			h:    []int{5, 12, 14, 21},
			s:    []int{6, 4, 7, 2},
			want: 23,
		},
		{
			name: "N=6",
			h:    []int{100, 100, 100, 100, 100, 1},
			s:    []int{1, 1, 1, 1, 1, 30},
			want: 105,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := ch06.Shooting(tt.h, tt.s); got != tt.want {
				t.Errorf("got %d, want %d\n", got, tt.want)
			}
		})
	}
}

func TestAscOrderIndices(t *testing.T) {
	t.Parallel()

	a := []int{12, 43, 7, 15, 9}
	want := []int{2, 4, 0, 3, 1}
	if got := ch06.AscOrderIndices(a); !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v\n", got, want)
	}
}
