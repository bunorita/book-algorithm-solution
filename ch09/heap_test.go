package ch09

import (
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	t.Parallel()

	h := NewIntMaxHeap()
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
		want []int
	}{
		{
			name: "case0", // 要素数 奇数
			a:    []int{5, 3, 1},
			want: []int{5, 3, 1},
		},
		{
			name: "case1", // 要素数 偶数
			a:    []int{5, 3, 7, 1},
			want: []int{7, 3, 5, 1},
		},
		{
			name: "case2",
			a:    []int{10, 0, 3, 8, 4, 1},
			want: []int{10, 8, 3, 0, 4, 1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h := NewIntMaxHeap()
			h.Push(tt.a...)

			if got := h.Array(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestHeapPop(t *testing.T) {
	t.Parallel()

	cond := func(p, c int) bool { return p >= c }

	tests := []*struct {
		name string
		h    *Heap[int]
		want []int
	}{
		{
			name: "case0",
			h:    &Heap[int]{arr: []int{5, 3, 1}, cond: cond},
			want: []int{3, 1},
		},
		{
			name: "case1",
			h:    &Heap[int]{arr: []int{7, 3, 5, 1}, cond: cond},
			want: []int{5, 3, 1},
		},
		{
			name: "case2",
			h:    &Heap[int]{arr: []int{10, 8, 3, 0, 4, 1}, cond: cond},
			want: []int{8, 4, 3, 0, 1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			popWant := tt.h.Array()[0]
			popGot, err := tt.h.Pop()
			if err != nil {
				t.Fatal(err)
			}
			if popGot != popWant {
				t.Errorf("got=%d, want=%d\n", popGot, popWant)
			}

			if got := tt.h.Array(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap(t *testing.T) {
	t.Parallel()

	h := NewIntMinHeap()
	h.Push(5, 3, 7, 1)

	got, err := h.Top()
	if err != nil {
		t.Fatal(err)
	}
	if want := 1; got != want {
		t.Errorf("got=%d, want=%d", got, want)
	}

	got, err = h.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if want := 1; got != want {
		t.Errorf("got=%d, want=%d", got, want)
	}

	got, err = h.Top()
	if err != nil {
		t.Fatal(err)
	}
	if want := 3; got != want {
		t.Errorf("got=%d, want=%d", got, want)
	}

	h.Push(2)
	got, err = h.Top()
	if err != nil {
		t.Fatal(err)
	}
	if want := 2; got != want {
		t.Errorf("got=%d, want=%d", got, want)
	}
}
