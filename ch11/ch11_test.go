package ch11_test

import (
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch11"
)

func TestBridge(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name  string
		n     int
		edges [][2]int
		want  int
	}{
		{
			name: "case1",
			n:    7,
			edges: [][2]int{
				{0, 2},
				{1, 6},
				{2, 3},
				{3, 4},
				{3, 5},
				{4, 5},
				{5, 6},
			},
			want: 4,
		},
		{
			name: "case2",
			n:    3,
			edges: [][2]int{
				{0, 1},
				{0, 2},
				{1, 2},
			},
			want: 0,
		},
		{
			name: "case3",
			n:    6,
			edges: [][2]int{
				{0, 1},
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch11.Bridge(tt.n, tt.edges)
			if got != tt.want {
				t.Errorf("got=%d, want=%d", got, tt.want)
			}
		})
	}
}
