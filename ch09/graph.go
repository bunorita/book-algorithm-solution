package ch09

import (
	"fmt"
	"sort"
)

type Graph [][]int

type Edge [2]int

// n = number of vertices
func NewGraph(n int, edges []Edge, directed bool) (*Graph, error) {
	g := make(Graph, n)
	for i := range g { // initialize
		g[i] = make([]int, 0)
	}

	for _, edge := range edges {
		for _, v := range edge {
			if v < 0 || v >= n {
				return nil, fmt.Errorf("unexpected vertex %d", v)
			}
		}
		a, b := edge[0], edge[1] // edge: a -> b
		g[a] = append(g[a], b)
		if !directed {
			g[b] = append(g[b], a)
		}
	}

	for _, adjacentList := range g {
		sort.Ints(adjacentList)
	}
	return &g, nil
}
