package ch14

import (
	"fmt"
	"sort"
)

func BellmanFord() []int {
	return []int{}
}

// directed & weighted graph
type Graph [][]Edge

type Edge struct {
	To int // 隣接頂点
	W  int
}

// edge [3]int: {from, to, weight}
func NewGraph(n int, edges [][3]int) (*Graph, error) {
	g := make(Graph, n)
	for i := range g { // initialize
		g[i] = make([]Edge, 0)
	}

	for _, edge := range edges {
		from, to, w := edge[0], edge[1], edge[2]
		for _, v := range []int{from, to} {
			if v < 0 || v >= n {
				return nil, fmt.Errorf("unexpected vertex %d", v)
			}
		}
		g[from] = append(g[from], Edge{To: to, W: w})
	}

	for _, edges := range g {
		sort.Slice(edges, func(i, j int) bool {
			return edges[i].To < edges[j].To
		})
	}

	return &g, nil
}
