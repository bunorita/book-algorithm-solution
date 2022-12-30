package ch08_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch08"
)

func TestNewArray(t *testing.T) {
	t.Parallel()

	want := &ch08.Array{1, 2, 3}
	got := ch08.NewArray(1, 2, 3)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestArrayGet(t *testing.T) {
	t.Parallel()

	arr := ch08.NewArray(1, 2, 3)
	i := 1
	want := 2

	if got := arr.Get(i); got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func TestArraySet(t *testing.T) {
	t.Parallel()

	arr := ch08.NewArray(1, 2, 3)
	arr.Set(1, 5)
	want := &ch08.Array{1, 5, 3}

	if got := arr; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func TestArrayAppend(t *testing.T) {
	t.Parallel()

	a := ch08.NewArray(1, 2, 3)
	a.Append(4)

	want := &ch08.Array{1, 2, 3, 4}
	if got := a; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestArrayInsert(t *testing.T) {
	t.Parallel()

	a := ch08.NewArray(1, 2, 3, 4)
	a.Insert(5, 2)
	want := &ch08.Array{1, 2, 5, 3, 4}

	if got := a; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestArrayRemove(t *testing.T) {
	t.Parallel()

	a := ch08.NewArray(1, 2, 3, 4)
	a.Remove(2)
	want := &ch08.Array{1, 3, 4}

	if got := a; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
