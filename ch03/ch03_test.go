package ch03_test

import (
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch03"
)

func TestLinearSearch(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name  string
		a     []int
		v     int
		exist bool
	}{
		{
			name:  "exist",
			a:     []int{4, 3, 12, 7, 11},
			v:     7,
			exist: true,
		},
		{
			name:  "not exist",
			a:     []int{4, 3, 12, 7, 11},
			v:     9,
			exist: false,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch03.LinearSearch(tt.a, tt.v)
			if got != tt.exist {
				t.Errorf("want %t, but got %t\n", tt.exist, got)
			}
		})
	}
}
