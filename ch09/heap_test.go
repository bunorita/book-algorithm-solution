package ch09_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch09"
)

func TestHeap(t *testing.T) {
	t.Parallel()

	h := ch09.NewHeap()
	h.Push(5, 3, 7, 1)

	got, err := h.Top()
	if err != nil {
		t.Fatal(err)
	}
	if want := 7; got != want {
		t.Errorf("got=%d, want=%d", got, want)
	}

	got, err = h.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if want := 7; got != want {
		t.Errorf("got=%d, want=%d", got, want)
	}

	got, err = h.Top()
	if err != nil {
		t.Fatal(err)
	}
	if want := 5; got != want {
		t.Errorf("got=%d, want=%d", got, want)
	}

	h.Push(11)
	got, err = h.Top()
	if err != nil {
		t.Fatal(err)
	}
	if want := 11; got != want {
		t.Errorf("got=%d, want=%d", got, want)
	}
}

func TestHeapPush(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name string
		a    []int
		want *ch09.Heap
	}{
		{
			name: "case0", // 要素数 奇数
			a:    []int{5, 3, 1},
			want: &ch09.Heap{5, 3, 1},
		},
		{
			name: "case1", // 要素数 偶数
			a:    []int{5, 3, 7, 1},
			want: &ch09.Heap{7, 3, 5, 1},
		},
		{
			name: "case2",
			a:    []int{10, 0, 3, 8, 4, 1},
			want: &ch09.Heap{10, 8, 3, 0, 4, 1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h := ch09.NewHeap()
			h.Push(tt.a...)

			if got := h; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestHeapPop(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name    string
		h, want *ch09.Heap
	}{
		{
			name: "case0",
			h:    &ch09.Heap{5, 3, 1},
			want: &ch09.Heap{3, 1},
		},
		{
			name: "case1",
			h:    &ch09.Heap{7, 3, 5, 1},
			want: &ch09.Heap{5, 3, 1},
		},
		{
			name: "case2",
			h:    &ch09.Heap{10, 8, 3, 0, 4, 1},
			want: &ch09.Heap{8, 4, 3, 0, 1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			popWant := (*tt.h)[0]
			popGot, err := tt.h.Pop()
			if err != nil {
				t.Fatal(err)
			}
			if popGot != popWant {
				t.Errorf("got=%d, want=%d\n", popGot, popWant)
			}

			if got := tt.h; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
