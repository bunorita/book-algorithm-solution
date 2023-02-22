package ch13

import (
	"github.com/bunorita/book-algorithm-solution/ch08"
	"github.com/bunorita/book-algorithm-solution/ch09"
)

func BFS(g *ch09.Graph, s int) ([]bool, error) {
	return search(g, s, ch08.NewQueue())
}

func DFS(g *ch09.Graph, s int) ([]bool, error) {
	return search(g, s, ch08.NewStack())
}

func dfsr(g *ch09.Graph, v int, seen *[]bool) {
	(*seen)[v] = true

	for _, next := range g.VerticesConnectedWith(v) {
		if (*seen)[next] {
			continue
		}
		dfsr(g, next, seen)
	}
}

type set interface {
	Push(x int) error
	Pop() (int, error)
	IsEmpty() bool
}

func search(g *ch09.Graph, s int, todo set) ([]bool, error) {
	seen := make([]bool, g.Size())

	// initialize
	seen[s] = true
	if err := todo.Push(s); err != nil {
		return nil, err
	}

	for !todo.IsEmpty() {
		v, err := todo.Pop()
		if err != nil {
			return nil, err
		}
		// vから辿れる頂点を全て調べる
		for _, x := range g.VerticesConnectedWith(v) {
			if seen[x] {
				continue
			}

			seen[x] = true
			if err := todo.Push(x); err != nil {
				return nil, err
			}
		}
	}

	return seen, nil
}

// recursive
func DFSr(g *ch09.Graph) []bool {
	n := g.Size()
	seen := make([]bool, n)
	for v := 0; v < n; v++ {
		dfsr(g, v, &seen)
	}
	return seen
}

// by BFS
func Distance(g *ch09.Graph, s int) ([]int, error) {
	dist := make([]int, g.Size())
	for i := range dist {
		dist[i] = -1 // 「未訪問」で初期化
	}
	dist[s] = 0

	todo := ch08.NewQueue()
	if err := todo.Push(s); err != nil {
		return nil, err
	}

	for !todo.IsEmpty() {
		v, err := todo.Pop()
		if err != nil {
			return nil, err
		}
		for _, x := range g.VerticesConnectedWith(v) {
			if dist[x] != -1 { // 探索済み
				continue
			}
			dist[x] = dist[v] + 1
			if err := todo.Push(x); err != nil {
				return nil, err
			}
		}

	}
	return dist, nil
}
