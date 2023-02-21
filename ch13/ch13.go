package ch13

import (
	"github.com/bunorita/book-algorithm-solution/ch08"
	"github.com/bunorita/book-algorithm-solution/ch09"
)

func Search(g *ch09.Graph, s int) ([]bool, error) {
	seen := make([]bool, g.Size())
	todo := ch08.NewQueue()
	// initialize
	seen[s] = true
	if err := todo.Enqueue(s); err != nil {
		return nil, err
	}

	for !todo.IsEmpty() {
		v, err := todo.Dequeue()
		if err != nil {
			return nil, err
		}
		// vから辿れる頂点を全て調べる
		for _, x := range g.VerticesConnectedWith(v) {
			if seen[x] {
				continue
			}

			seen[x] = true
			if err := todo.Enqueue(x); err != nil {
				return nil, err
			}
		}
	}

	return seen, nil
}
