package ch14

import (
	"errors"
	"fmt"
	"math"
	"sort"

	"github.com/bunorita/book-algorithm-solution/ch09"
	"github.com/bunorita/book-algorithm-solution/intutil"
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

// 14.4
func DijkstraByHeap(g *Graph, s int) ([]int, error) {
	dist := make([]int, len(*g))
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[s] = 0

	type pair struct{ v, d int } // {v, dist[v]}
	// Min Heap. dist[v]が最小のものを取り出せる
	h := ch09.NewHeap(func(p, c pair) bool { return p.d <= c.d })
	h.Push(pair{v: 0, d: dist[0]})

	for !h.IsEmpty() {
		// v: 使用済みでない頂点のうちdist[v]が最小の頂点
		// d: vに対するキー値 dist[v]
		top, err := h.Pop()
		if err != nil {
			return nil, err
		}
		v, d := top.v, top.d
		fmt.Printf("v: %d, d: %d\n", v, d)
		if d > dist[v] { // 複数回緩和で最小でなくなった値dは無視
			continue
		}

		for _, edge := range (*g)[v] {
			to := edge.To
			if chmin(&dist[to], dist[v]+edge.W) {
				fmt.Printf("    to: %d, dist: %d\n", to, dist[to])
				h.Push(pair{v: to, d: dist[to]})
			}

		}
	}

	return dist, nil
}

// 14.5 Floyd-Warshall algorithm
// https://atcoder.jp/contests/abc208/tasks/abc208_d
func ShortestPathQueries2(n int, edges [][3]int) (int, error) {
	dp := make([][]int, n)
	const inf = math.MaxInt
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	// dp初期条件
	for _, edge := range edges {
		a, b, w := edge[0], edge[1], edge[2]
		dp[a][b] = w
	}
	for v := 0; v < n; v++ {
		dp[v][v] = 0
	}

	var result int
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				chmin(&dp[i][j], intutil.SafeAdd(dp[i][k], dp[k][j]))
				if dp[i][j] < inf {
					result += dp[i][j]
				}
			}
		}
	}

	for v := 0; v < n; v++ {
		if dp[v][v] < 0 {
			return 0, errors.New("negative cycle exists")
		}
	}

	return result, nil
}
