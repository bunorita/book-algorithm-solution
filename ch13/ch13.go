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
