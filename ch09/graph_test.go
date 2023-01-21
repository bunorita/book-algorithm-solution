package ch09_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch09"
)

func TestNewGraph(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name     string
		n        int
		edges    []ch09.Edge
		directed bool
		want     *ch09.Graph
	}{
		{
			name:     "n=2 undirected",
			n:        2,
			edges:    []ch09.Edge{{0, 1}},
			directed: false,
			want: &ch09.Graph{
				{1},
				{0},
			},
		},
		{
			name:     "n=2 directed",
			n:        2,
			edges:    []ch09.Edge{{0, 1}},
			directed: true,
			want: &ch09.Graph{
				{1},
				{},
			},
		},
		{
			name: "n=8 undirected",
			n:    8,
			edges: []ch09.Edge{
				{4, 1},
				{4, 2},
				{4, 6},
				{1, 3},
				{1, 6},
				{2, 5},
				{2, 7},
				{6, 7},
				{3, 0},
				{3, 7},
				{7, 0},
				{0, 5},
			},
			directed: false,
			want: &ch09.Graph{
				{3, 5, 7},
				{3, 4, 6},
				{4, 5, 7},
				{0, 1, 7},
				{1, 2, 6},
				{0, 2},
				{1, 4, 7},
				{0, 2, 3, 6},
			},
		},
		{
			name: "n=8 directed",
			n:    8,
			edges: []ch09.Edge{
				{4, 1},
				{4, 2},
				{4, 6},
				{1, 3},
				{1, 6},
				{2, 5},
				{2, 7},
				{6, 7},
				{3, 0},
				{3, 7},
				{7, 0},
				{0, 5},
			},
			directed: true,
			want: &ch09.Graph{
				{5},
				{3, 6},
				{5, 7},
				{0, 7},
				{1, 2, 6},
				{},
				{7},
				{0},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ch09.NewGraph(tt.n, tt.edges, tt.directed)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}

		})
	}
}
