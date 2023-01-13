package ch08_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch08"
)

func TestNewStack(t *testing.T) {
	t.Parallel()

	s, err := ch08.NewStack(3, 7, 5, 4)
	if err != nil {
		t.Fatal(err)
	}
	got := s.Values()
	want := []int{3, 7, 5, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestStackPop(t *testing.T) {
	t.Parallel()

	s, err := ch08.NewStack(3, 7, 5, 4)
	if err != nil {
		t.Fatal(err)
	}

	want := 4
	got, err := s.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}

	gotValues := s.Values()
	wantValues := []int{3, 7, 5}
	if !reflect.DeepEqual(gotValues, wantValues) {
		t.Errorf("values got: %v, want: %v", gotValues, wantValues)
	}
}
