package ch08_test

import (
	"strings"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch08"
)

func TestRPN(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		expr []string
		want int
	}{
		{
			expr: []string{"3", "4", "+"},
			want: 7,
		},
		{
			expr: []string{"15", "3", "/"},
			want: 5,
		},
		{
			expr: []string{"3", "4", "+", "1", "2", "-", "*"},
			want: -7,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(strings.Join(tt.expr, " "), func(t *testing.T) {
			t.Parallel()

			got, err := ch08.RPN(tt.expr)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("got=%d, want=%d", got, tt.want)
			}
		})
	}
}
