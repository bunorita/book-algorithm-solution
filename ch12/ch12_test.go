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

func TestBucketSort(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name string
		a    []int
		want []int
	}{
		{
			name: "case0",
			a:    []int{3, 1, 2},
			want: []int{1, 2, 3},
		},
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
		{
			name: "case3", // 値重複あり
			a:    []int{10, 12, 12, 1, 4, 17, 4, 1},
			want: []int{1, 1, 4, 4, 10, 12, 12, 17},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ch12.BucketSort(&tt.a)
			if got := tt.a; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestAscOrderIndices(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name string
		// aは値重複なし
		a, want []int
	}{
		{
			name: "case0",
			a:    []int{3, 1, 2},
			want: []int{2, 0, 1},
		},
		{
			name: "case1",
			a:    []int{5, 3, 7, 1},
			want: []int{2, 1, 3, 0},
		},
		{
			name: "case2",
			a:    []int{10, 12, 15, 3, 8, 17, 4, 1},
			want: []int{4, 5, 6, 1, 3, 7, 2, 0},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch12.AscOrderIndices(tt.a)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestEnergyDrinkCollector(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name   string
		stores [][2]int
		m      int
		want   int
	}{
		{
			name: "case1",
			stores: [][2]int{
				{4, 9},
				{2, 4},
			},
			m:    5,
			want: 12,
		},
		{
			name: "case2",
			stores: [][2]int{
				{6, 18},
				{2, 5},
				{3, 10},
				{7, 9},
			},
			m:    30,
			want: 130,
		},
		{
			name: "case3",
			stores: [][2]int{
				{1_000_000_000, 100_000},
			},
			m:    100_000,
			want: 100_000_000_000_000,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := ch12.EnergyDrinkCollector(tt.stores, tt.m); got != tt.want {
				t.Errorf("got=%d, want=%d", got, tt.want)
			}
		})
	}
}
