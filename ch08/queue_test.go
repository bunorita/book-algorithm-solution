package ch08_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch08"
)

func TestQueue(t *testing.T) {
	t.Parallel()

	// initialize
	q, err := ch08.NewQueue(3, 5, 7)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := q.Values(), []int{3, 5, 7}; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	// dequeue
	got, err := q.Dequeue()
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
	if got, want := q.Values(), []int{5, 7}; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	// enquue
	q.Enqueue(9)
	if got, want := q.Values(), []int{5, 7, 9}; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
