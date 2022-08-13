package ch04_test

import (
	"fmt"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch04"
)

func TestSumUpTo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		upto int
		want int
	}{
		{
			upto: 0,
			want: 0,
		},
		{
			upto: 10,
			want: 55,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("sum up to %d", tt.upto), func(t *testing.T) {
			t.Parallel()

			got := ch04.SumUpTo(tt.upto)
			if got != tt.want {
				t.Errorf("want %d, but got %d\n", tt.want, got)
			}
		})
	}
}

func TestGCD(t *testing.T) {
	t.Parallel()

	m, n := 51, 15
	want := 3
	got := ch04.GCD(m, n)
	if got != want {
		t.Errorf("want %d, but got %d\n", want, got)
	}
}

func TestFib(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n    int
		want int
	}{
		{
			n:    0,
			want: 0,
		},
		{
			n:    1,
			want: 1,
		},
		{
			n:    12,
			want: 144,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			t.Parallel()

			got := ch04.Fib(tt.n)
			if got != tt.want {
				t.Errorf("Fib() want %d, but got %d\n", tt.want, got)
			}

			gotLoop := ch04.FibLoop(tt.n)
			if gotLoop != tt.want {
				t.Errorf("FibLoop() want %d, but got %d\n", tt.want, gotLoop)
			}

			gotMemo := ch04.FibMemo(tt.n)
			if gotMemo != tt.want {
				t.Errorf("FibMemo() want %d, but got %d\n", tt.want, gotMemo)
			}
		})
	}
}

func BenchmarkFib_40(b *testing.B) {
	benchmarkFib(b, ch04.Fib, 40)
}

func BenchmarkFibLoop_40(b *testing.B) {
	benchmarkFib(b, ch04.FibLoop, 40)
}

func BenchmarkFibMemo_40(b *testing.B) {
	benchmarkFib(b, ch04.FibMemo, 40)
}

func benchmarkFib(b *testing.B, fib func(int) int, n int) {
	for i := 0; i < b.N; i++ {
		fib(n)
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
		{
			name: "true2",
			a:    []int{3, 2, 6, 5},
			w:    14,
			want: true,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch04.PartialSumEquals(tt.a, tt.w)
			if got != tt.want {
				t.Errorf("want %t, but got %t\n", tt.want, got)
			}
		})
	}
}
