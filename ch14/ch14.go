package ch14

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

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

func BellmanFord(g *Graph, s int) ([]int, error) {
	n := len(*g)
	dist := make([]int, n)
	for i := range dist { // initialize
		dist[i] = math.MaxInt
	}
	dist[s] = 0

	for i := 0; i < n; i++ {
		fmt.Printf("iteration %d\n", i)
		var updated bool
		for v := 0; v < n; v++ {
			if dist[v] == math.MaxInt { // vは始点sから到達できない
				continue
			}
			for _, edge := range (*g)[v] {
				if chmin(&dist[edge.To], dist[v]+edge.W) {
					updated = true
					fmt.Printf("	relaxation (%d => %d) %d\n", v, edge.To, dist[edge.To])
				}
			}
		}

		if !updated { // 更新がなければ既に最短路が求められている
			break
		}
		if i == n-1 { // n回目反復で更新があれば負閉路を持つ
			return nil, errors.New("negative cycle exists")
		}
	}
	return dist, nil
}

func chmin(a *int, b int) bool {
	if (*a) > b {
		*a = b
		return true
	}
	return false
}

// 14.3
func Dijkstra(g *Graph, s int) []int {
	n := len(*g)
	used := make([]bool, n)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[s] = 0
	for i := 0; i < n; i++ {
		fmt.Printf("iteration %d\n", i)
		// 未使用の頂点の内、dist値が最小の頂点を探す
		minDist := math.MaxInt
		minDistV := -1
		for v := 0; v < n; v++ {
			if !used[v] && dist[v] < minDist {
				minDist = dist[v]
				minDistV = v
			}
		}
		// そのような頂点がなければ終了
		if minDistV == -1 {
			break
		}

		// minDistVを始点として各辺を緩和する
		for _, edge := range (*g)[minDistV] {
			chmin(&dist[edge.To], dist[minDistV]+edge.W)
			fmt.Printf("	relaxation (%d => %d) %d\n", minDistV, edge.To, dist[edge.To])
		}
		used[minDistV] = true
	}
	return dist
}
