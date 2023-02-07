package ch12_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch12"
)

func TestInsertionSort(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name string
		a    []int
		want []int
	}{
		{
			name: "case1",
			a:    []int{4, 1},
			want: []int{1, 4},
		},
		{
			name: "case2",
			a:    []int{4, 1, 3},
			want: []int{1, 3, 4},
		},
		{
			name: "case3",
			a:    []int{4, 1, 3, 5, 2},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch12.InsertionSort(tt.a)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name string
		a    []int
		want []int
	}{
		{
			name: "case1",
			a:    []int{4, 2, 3, 1},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "case2",
			a:    []int{12, 9, 15, 3, 8, 17, 6, 1},
			want: []int{1, 3, 6, 8, 9, 12, 15, 17},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch12.MergeSort(tt.a)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
