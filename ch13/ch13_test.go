package ch13_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch09"
	"github.com/bunorita/book-algorithm-solution/ch13"
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
				t.Errorf("got: %t, want: %t", got, tt.want)
			}
		})
	}
}
