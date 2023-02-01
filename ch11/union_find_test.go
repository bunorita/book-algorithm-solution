package ch11_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch11"
)

func TestUnionFind(t *testing.T) {
	t.Parallel()

	uf := ch11.NewUnionFind(7)

	if got, want := uf.Unions(), [][]int{{0}, {1}, {2}, {3}, {4}, {5}, {6}}; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	uf.Unite(1, 2)

	if got, want := uf.Unions(), [][]int{{0}, {1, 2}, {3}, {4}, {5}, {6}}; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	uf.Unite(2, 3)

	if got, want := uf.Unions(), [][]int{{0}, {1, 2, 3}, {4}, {5}, {6}}; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	uf.Unite(5, 6)

	if got, want := uf.Unions(), [][]int{{0}, {1, 2, 3}, {4}, {5, 6}}; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	if got, want := uf.IsSame(1, 3), true; got != want {
		t.Errorf("got: %t, want: %t", got, want)
	}
	if got, want := uf.IsSame(2, 5), false; got != want {
		t.Errorf("got: %t, want: %t", got, want)
	}

	uf.Unite(1, 6)

	if got, want := uf.Unions(), [][]int{{0}, {1, 2, 3, 5, 6}, {4}}; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	if got, want := uf.IsSame(2, 5), true; got != want {
		t.Errorf("got: %t, want: %t", got, want)
	}
}
