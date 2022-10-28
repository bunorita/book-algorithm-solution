package ch05_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch05"
)

func TestFrog1(t *testing.T) {
	t.Parallel()

	h := []int{2, 9, 4, 5, 1, 6, 10}
	want := []int{0, 7, 2, 3, 5, 4, 8}

	tests := []*struct {
		name  string
		frog1 func([]int) []int
	}{
		{
			name:  "pull-based",
			frog1: ch05.Frog1,
		},
		{
			name:  "relaxation",
			frog1: ch05.Frog1Relaxation,
		},
		{
			name:  "push-based",
			frog1: ch05.Frog1PushBased,
		},
		{
			name:  "recursive",
			frog1: ch05.Frog1R,
		},
		{
			name:  "recursive & memoization",
			frog1: ch05.Frog1RM,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.frog1(h)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

var h = make([]int, 20)

func init() {
	for i := 0; i < 20; i++ {
		h = append(h, i)
	}
}

func BenchmarkFrog1R_20(b *testing.B) {
	benchmarkFrog1R(b, ch05.Frog1R, h)
}

func BenchmarkFrog1RM_20(b *testing.B) {
	benchmarkFrog1R(b, ch05.Frog1RM, h)
}

func benchmarkFrog1R(b *testing.B, frog1 func([]int) []int, h []int) {
	for i := 0; i < b.N; i++ {
		frog1(h)
	}
}

func TestKnapsack(t *testing.T) {
	t.Parallel()

	items := []*ch05.Item{
		{2, 3},
		{1, 2},
		{3, 6},
		{2, 1},
		{1, 3},
		{5, 85},
	}
	wmax := 9
	want := 94
	if got := ch05.Knapsack(items, wmax); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestEditDist(t *testing.T) {
	t.Parallel()

	s1 := "logistic"
	s2 := "algorithm"
	want := 6
	if got := ch05.EditDist(s1, s2); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestBoxingOranges(t *testing.T) {
	// 10 oranges
	oranges := []int{3, 5, 4, 9, 11, 8, 10, 1, 7, 8} // sizes of each orange

	want := 40
	if got := ch05.BoxingOranges(oranges); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestVacation(t *testing.T) {
	t.Parallel()

	a := []int{5, 5, 2, 4, 8}
	b := []int{1, 3, 7, 9, 1}
	c := []int{4, 4, 3, 6, 3}
	x := make([][3]int, len(a))
	for i := range a {
		x[i] = [3]int{a[i], b[i], c[i]}
	}

	tests := []*struct {
		n    int
		want int
	}{
		{
			n:    1,
			want: 5,
		},
		{
			n:    2,
			want: 9,
		},
		{
			n:    3,
			want: 16,
		},
		{
			n:    4,
			want: 22,
		},
		{
			n:    5,
			want: 30,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			t.Parallel()

			if got := ch05.Vacation(x[:tt.n]); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestPartialSumEquals(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		a    []int
		w    int
		want bool
	}{
		{
			name: "true",
			a:    []int{1, 2, 4, 5, 11},
			w:    10,
			want: true,
		},
		{
			name: "false",
			a:    []int{1, 5, 8, 11},
			w:    10,
			want: false,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch05.PartialSumEquals(tt.a, tt.w)
			if got != tt.want {
				t.Errorf("want %t, but got %t\n", tt.want, got)
			}
		})
	}
}

func TestCountPartialSumLTE(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		a    []int
		w    int
		want int
	}{
		{
			name: "1",
			a:    []int{1, 2, 4},
			w:    10,
			want: 7, // 1,2,3,4,5,6,7
		},
		{
			name: "2",
			a:    []int{1, 2, 4, 5},
			w:    10,
			want: 10, // 1,2,3,4,5,6,7,8,9,10

		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch05.CountPartialSumLTE(tt.a, tt.w)
			if got != tt.want {
				t.Errorf("want %d, but got %d\n", tt.want, got)
			}
		})
	}
}

func TestKIntsOrLessPartialSumEquals(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		a    []int
		k    int
		w    int
		want bool
	}{
		{
			name: "true",
			a:    []int{1, 2, 4, 5, 11},
			k:    3,
			w:    10,
			want: true,
		},
		{
			name: "false",
			a:    []int{1, 2, 4, 5, 11},
			k:    2,
			w:    10,
			want: false,
		},
		{
			name: "debug",
			a:    []int{1, 2},
			k:    2,
			w:    3,
			want: true,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := ch05.KIntsOrLessPartialSumEquals(tt.a, tt.k, tt.w); got != tt.want {
				t.Errorf("KIntsOrLessPartialSumEquals want %t, but got %t\n", tt.want, got)
			}

			if got := ch05.KIntsOrLessPartialSumEquals2(tt.a, tt.k, tt.w); got != tt.want {
				t.Errorf("KIntsOrLessPartialSumEquals2 want %t, but got %t\n", tt.want, got)
			}
		})
	}
}
