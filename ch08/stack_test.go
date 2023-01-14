package ch08_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch08"
)

func TestStack(t *testing.T) {
	t.Parallel()

	// initialize
	s, err := ch08.NewStack(3, 7, 5, 4)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := s.Values(), []int{3, 7, 5, 4}; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	// pop
	got, err := s.Pop()
	if err != nil {
		t.Fatal(err)
	}
	want := 4
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
	if got, want := s.Values(), []int{3, 7, 5}; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	// push
	if err := s.Push(9); err != nil {
		t.Fatal(err)
	}
	if got, want := s.Values(), []int{3, 7, 5, 9}; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
