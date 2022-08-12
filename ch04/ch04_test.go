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
		})
	}
}
