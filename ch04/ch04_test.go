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
