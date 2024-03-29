package ch13_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch09"
	"github.com/bunorita/book-algorithm-solution/ch13"
	"github.com/bunorita/book-algorithm-solution/util"
)

func TestBFS(t *testing.T) {
	t.Parallel()

	n := 9
	g, err := ch09.NewGraph(n, []ch09.Edge{
		{0, 1},
		{0, 4},
		{0, 2},
		{1, 3},
		{1, 8},
		{3, 8},
		{4, 8},
		{2, 5},
		{5, 8},
		{3, 7},
		{5, 6},
		{6, 7},
	}, false)
	if err != nil {
		t.Fatal(err)
	}

	want := make([]bool, n)
	for i := range want {
		want[i] = true
	}

	got, err := ch13.BFS(g, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v\n", got, want)
	}

}

func TestDFS(t *testing.T) {
	t.Parallel()

	n := 9
	g, err := ch09.NewGraph(n, []ch09.Edge{
		{0, 1},
		{0, 4},
		{0, 2},
		{1, 3},
		{1, 8},
		{3, 8},
		{4, 8},
		{2, 5},
		{5, 8},
		{3, 7},
		{5, 6},
		{6, 7},
	}, false)
	if err != nil {
		t.Fatal(err)
	}

	want := make([]bool, n)
	for i := range want {
		want[i] = true
	}

	got, err := ch13.DFS(g, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v\n", got, want)
	}

}

func TestDFSr(t *testing.T) {
	t.Parallel()

	n := 8
	g, err := ch09.NewGraph(n, []ch09.Edge{
		{0, 5},
		{1, 3},
		{1, 6},
		{2, 5},
		{2, 7},
		{3, 0},
		{3, 7},
		{4, 1},
		{4, 2},
		{4, 6},
		{6, 7},
		{7, 0},
	}, true)
	if err != nil {
		t.Fatal(err)
	}

	want := make([]bool, n)
	for i := range want {
		want[i] = true
	}

	got := ch13.DFSr(g)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v\n", got, want)
	}
}

func TestDistance(t *testing.T) {
	t.Parallel()

	n := 9
	g, err := ch09.NewGraph(n, []ch09.Edge{
		{0, 1},
		{0, 4},
		{0, 2},
		{1, 3},
		{1, 8},
		{3, 8},
		{4, 8},
		{2, 5},
		{5, 8},
		{3, 7},
		{5, 6},
		{6, 7},
	}, false)
	if err != nil {
		t.Fatal(err)
	}

	want := []int{0, 1, 1, 2, 1, 2, 3, 3, 2}

	got, err := ch13.Distance(g, 0)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v\n", got, want)
	}
}

func TestPathExists(t *testing.T) {
	t.Parallel()

	// p158 図10.4
	n := 8
	g, err := ch09.NewGraph(n, []ch09.Edge{
		{0, 5},
		{1, 3},
		{1, 6},
		{2, 5},
		{2, 7},
		{3, 0},
		{3, 7},
		{4, 1},
		{4, 2},
		{4, 6},
		{6, 7},
		{7, 0},
	}, true)
	if err != nil {
		t.Fatal(err)
	}

	tests := []*struct {
		name string
		g    *ch09.Graph
		s, t int
		want bool
	}{
		{
			name: "case1",
			g:    g,
			s:    6,
			t:    5,
			want: true,
		},
		{
			name: "case2",
			g:    g,
			s:    3,
			t:    6,
			want: false,
		},
		{
			name: "case3",
			g:    g,
			s:    4,
			t:    0,
			want: true,
		},
		{
			name: "case4",
			g:    g,
			s:    0,
			t:    4,
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := ch13.PathExists(tt.g, tt.s, tt.t); got != tt.want {
				t.Errorf("DFS got: %t, want: %t", got, tt.want)
			}

			got, err := ch13.PathExistsBFS(tt.g, tt.s, tt.t)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("BFS got: %t, want: %t", got, tt.want)
			}
		})
	}
}

func TestIsBipartite(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name      string
		graphSize int
		edges     []ch09.Edge
		want      bool
	}{
		{
			name:      "case1",
			graphSize: 5,
			edges: []ch09.Edge{
				{0, 1},
				{0, 3},
				{1, 2},
				{1, 4},
				{3, 4},
			},
			want: true,
		},
		{
			name:      "case2",
			graphSize: 5,
			edges: []ch09.Edge{
				{0, 1},
				{0, 3},
				{1, 2},
				{1, 4},
				{3, 4},
				{2, 4},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			g, err := ch09.NewGraph(tt.graphSize, tt.edges, false)
			if err != nil {
				t.Fatal(err)
			}

			if got := ch13.IsBipartite(g); got != tt.want {
				t.Errorf("DFS got: %t, want: %t", got, tt.want)
			}

			got, err := ch13.IsBipartiteBFS(g)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("BFS got: %t, want: %t", got, tt.want)
			}
		})
	}
}

func TestTopologicalSort(t *testing.T) {
	t.Parallel()

	n := 8
	g, err := ch09.NewGraph(n, []ch09.Edge{
		{0, 5},
		{1, 3},
		{1, 6},
		{2, 5},
		{2, 7},
		{3, 0},
		{3, 7},
		{4, 1},
		{4, 2},
		{4, 6},
		{6, 7},
		{7, 0},
	}, true)
	if err != nil {
		t.Fatal(err)
	}
	want := []int{4, 2, 1, 6, 3, 7, 0, 5}
	got := ch13.TopologicalSort(g)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestTreeSearchDepth(t *testing.T) {
	t.Parallel()

	tree, err := ch09.NewGraph(15, []ch09.Edge{
		{0, 1},
		{0, 4},
		{0, 11},
		{1, 2},
		{1, 3},
		{4, 5},
		{4, 8},
		{5, 6},
		{5, 7},
		{8, 9},
		{8, 10},
		{11, 12},
		{11, 13},
		{13, 14},
	}, false)
	if err != nil {
		t.Fatal(err)
	}

	want := []int{0, 1, 2, 2, 1, 2, 3, 3, 2, 3, 3, 1, 2, 2, 3}
	got := ch13.TreeSearchDepth(tree)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestTreeSearchSubtreeSize(t *testing.T) {
	t.Parallel()

	tree, err := ch09.NewGraph(15, []ch09.Edge{
		{0, 1},
		{0, 4},
		{0, 11},
		{1, 2},
		{1, 3},
		{4, 5},
		{4, 8},
		{5, 6},
		{5, 7},
		{8, 9},
		{8, 10},
		{11, 12},
		{11, 13},
		{13, 14},
	}, false)
	if err != nil {
		t.Fatal(err)
	}

	want := []int{15, 3, 1, 1, 7, 3, 1, 1, 3, 1, 1, 4, 1, 2, 1}
	got := ch13.TreeSearchSubtreeSize(tree)
	util.TestDiff(t, got, want)
}

func TestCountConnectedGraph(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name  string
		n     int
		edges []ch09.Edge
		want  int
	}{
		{
			name:  "case1",
			n:     1,
			edges: []ch09.Edge{},
			want:  1,
		},
		{
			name:  "case2",
			n:     2,
			edges: []ch09.Edge{},
			want:  2,
		},
		{
			name: "case3",
			n:    2,
			edges: []ch09.Edge{
				{0, 1},
			},
			want: 1,
		},
		{
			name: "case4",
			n:    7,
			edges: []ch09.Edge{
				{1, 2},
				{1, 3},
				{5, 6},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			g, err := ch09.NewGraph(tt.n, tt.edges, false)
			if err != nil {
				t.Fatal(err)
			}
			if got := ch13.CountConnectedGraph(g); got != tt.want {
				t.Errorf("got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func TestShortestPathOfMaze(t *testing.T) {
	t.Parallel()

	maze := `
.#....#G
.#.#....
...#.##.
#.##...#
...###.#
.#.....#
...#.#..
S.......
`
	// o -- y (w)
	// |
	// |
	// x (h)

	want := 16 // steps
	got, err := ch13.ShortestPathOfMaze(8, 8, maze)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}

}
