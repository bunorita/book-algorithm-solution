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
