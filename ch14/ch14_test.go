package ch14_test

import (
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch14"
	"github.com/bunorita/book-algorithm-solution/util"
)

func TestNewGraph(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name  string
		n     int
		edges [][3]int
		want  *ch14.Graph
	}{
		{
			name: "has negative weight edges",
			n:    6,
			edges: [][3]int{
				{0, 1, 3},
				{0, 3, 100},
				{1, 2, 50},
				{1, 3, 57},
				{1, 4, -4},
				{2, 3, -10},
				{2, 5, 100},
				{3, 1, -5},
				{4, 2, 57},
				{4, 3, 25},
				{4, 5, 8},
			},
			want: &ch14.Graph{
				[]ch14.Edge{
					{To: 1, W: 3},
					{To: 3, W: 100},
				},
				[]ch14.Edge{
					{To: 2, W: 50},
					{To: 3, W: 57},
					{To: 4, W: -4},
				},
				[]ch14.Edge{
					{To: 3, W: -10},
					{To: 5, W: 100},
				},
				[]ch14.Edge{
					{To: 1, W: -5},
				},
				[]ch14.Edge{
					{To: 2, W: 57},
					{To: 3, W: 25},
					{To: 5, W: 8},
				},
				[]ch14.Edge{},
			},
		},
		{
			name: "no negative weight edges",
			n:    6,
			edges: [][3]int{
				{0, 1, 3},
				{0, 2, 5},
				{1, 2, 4},
				{1, 3, 12},
				{2, 3, 9},
				{2, 4, 4},
				{3, 5, 2},
				{4, 3, 7},
				{4, 5, 8},
			},
			want: &ch14.Graph{
				[]ch14.Edge{
					{To: 1, W: 3},
					{To: 2, W: 5},
				},
				[]ch14.Edge{
					{To: 2, W: 4},
					{To: 3, W: 12},
				},
				[]ch14.Edge{
					{To: 3, W: 9},
					{To: 4, W: 4},
				},
				[]ch14.Edge{
					{To: 5, W: 2},
				},
				[]ch14.Edge{
					{To: 3, W: 7},
					{To: 5, W: 8},
				},
				[]ch14.Edge{},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ch14.NewGraph(tt.n, tt.edges)
			if err != nil {
				t.Fatal(err)
			}
			util.TestDiff(t, got, tt.want)
		})
	}
}

func TestBellmanFord(t *testing.T) {
	t.Parallel()

	g, err := ch14.NewGraph(6, [][3]int{
		{0, 1, 3},
		{0, 3, 100},
		{1, 2, 50},
		{1, 3, 57},
		{1, 4, -4},
		{2, 3, -10},
		{2, 5, 100},
		{3, 1, -5},
		{4, 2, 57},
		{4, 3, 25},
		{4, 5, 8},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := []int{0, 3, 53, 24, -1, 7}
	got, err := ch14.BellmanFord(g, 0)
	if err != nil {
		t.Fatal(err)
	}
	util.TestDiff(t, got, want)
}

func TestDijkstra(t *testing.T) {
	t.Parallel()

	g, err := ch14.NewGraph(6, [][3]int{
		{0, 1, 3},
		{0, 2, 5},
		{1, 2, 4},
		{1, 3, 12},
		{2, 3, 9},
		{2, 4, 4},
		{3, 5, 2},
		{4, 3, 7},
		{4, 5, 8},
	})
	if err != nil {
		t.Fatal(err)
	}
	want := []int{0, 3, 5, 14, 9, 16}

	t.Run("simple", func(t *testing.T) {
		t.Parallel()

		got := ch14.Dijkstra(g, 0)
		util.TestDiff(t, got, want)
	})

	t.Run("heap", func(t *testing.T) {
		t.Parallel()

		got, err := ch14.DijkstraByHeap(g, 0)
		if err != nil {
			t.Fatal(err)
		}
		util.TestDiff(t, got, want)
	})
}

func TestShortestPathQueries2(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name  string
		n     int
		edges [][3]int
		want  int
	}{
		{
			name: "case1",
			n:    3,
			edges: [][3]int{
				{1, 2, 3},
				{2, 3, 2},
			},
			want: 25,
		},
		{
			name:  "case2",
			n:     3,
			edges: [][3]int{},
			want:  0,
		},
		{
			name: "case3",
			n:    5,
			edges: [][3]int{
				{1, 2, 6},
				{1, 3, 10},
				{1, 4, 4},
				{1, 5, 1},
				{2, 1, 5},
				{2, 3, 9},
				{2, 4, 8},
				{2, 5, 6},
				{3, 1, 5},
				{3, 2, 1},
				{3, 4, 7},
				{3, 5, 9},
				{4, 1, 4},
				{4, 2, 6},
				{4, 3, 4},
				{4, 5, 8},
				{5, 1, 2},
				{5, 2, 5},
				{5, 3, 6},
				{5, 4, 5},
			},
			want: 517,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// 1-based => 0-based indexing
			for i := range tt.edges {
				tt.edges[i][0]--
				tt.edges[i][1]--
			}

			got, err := ch14.ShortestPathQueries2(tt.n, tt.edges)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("got: %d, want: %d", got, tt.want)
			}
		})
	}
}
