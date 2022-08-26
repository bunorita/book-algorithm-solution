package ch05_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch05"
)

func TestFrog1(t *testing.T) {
	t.Parallel()

	h := []int{2, 9, 4, 5, 1, 6, 10}
	want := []int{0, 7, 2, 3, 5, 4, 8}
	got := ch05.Frog1(h)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but want %v", got, want)
	}
}

func TestFrog1Relaxation(t *testing.T) {
	t.Parallel()

	h := []int{2, 9, 4, 5, 1, 6, 10}
	want := []int{0, 7, 2, 3, 5, 4, 8}

	got := ch05.Frog1Relaxation(h)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but want %v", got, want)
	}
}
