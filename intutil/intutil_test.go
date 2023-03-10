package intutil_test

import (
	"math"
	"testing"

	"github.com/bunorita/book-algorithm-solution/intutil"
)

func TestSafeAdd(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name       string
		a, b, want int
	}{
		{name: "case0", a: 1, b: 2, want: 3},

		{name: "max1", a: math.MaxInt, b: 1, want: math.MaxInt},
		{name: "max2", a: math.MaxInt - 1, b: 2, want: math.MaxInt},
		{name: "max3", a: math.MaxInt - 2, b: 1, want: math.MaxInt - 1},

		{name: "min1", a: math.MinInt, b: -1, want: math.MinInt},
		{name: "min2", a: math.MinInt + 1, b: -2, want: math.MinInt},
		{name: "min3", a: math.MinInt + 2, b: -1, want: math.MinInt + 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := intutil.SafeAdd(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got: %d, want: %d", got, tt.want)
			}
		})
	}
}
