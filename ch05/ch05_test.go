package ch05_test

import (
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
