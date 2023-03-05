package ch14_test

import (
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch14"
	"github.com/bunorita/book-algorithm-solution/util"
)

func TestNewGraph(t *testing.T) {
	t.Parallel()

	n := 6
	edges := [][3]int{
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
	}

	want := &ch14.Graph{
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
	}

	got, err := ch14.NewGraph(n, edges)
	if err != nil {
		t.Fatal(err)
	}
	util.TestDiff(t, got, want)
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
