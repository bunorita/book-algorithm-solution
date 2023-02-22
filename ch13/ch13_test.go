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
