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

			ch12.InsertionSort(&tt.a)
			if got := tt.a; !reflect.DeepEqual(got, tt.want) {
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

			ch12.MergeSort(&tt.a)
			if got := tt.a; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name string
		a    []int
		want []int
	}{
		{
			name: "case1",
			a:    []int{10, 12, 15, 3, 8, 17, 4, 1},
			want: []int{1, 3, 4, 8, 10, 12, 15, 17},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ch12.QuickSort(&tt.a)
			if got := tt.a; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name string
		a    []int
		want []int
	}{
		{
			name: "case1",
			a:    []int{5, 3, 7, 1},
			want: []int{1, 3, 5, 7},
		},
		{
			name: "case2",
			a:    []int{10, 12, 15, 3, 8, 17, 4, 1},
			want: []int{1, 3, 4, 8, 10, 12, 15, 17},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ch12.HeapSort(&tt.a)
			if got := tt.a; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestHeapSort2(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name string
		a    []int
		want []int
	}{
		{
			name: "case1",
			a:    []int{5, 3, 7, 1},
			want: []int{1, 3, 5, 7},
		},
		{
			name: "case2",
			a:    []int{10, 12, 15, 3, 8, 17, 4, 1},
			want: []int{1, 3, 4, 8, 10, 12, 15, 17},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ch12.HeapSort2(&tt.a)
			if got := tt.a; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
