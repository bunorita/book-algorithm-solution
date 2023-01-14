package ch08_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch08"
)

func TestStack(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name  string
		stack ch08.IStack
	}{
		{name: "by array", stack: ch08.NewStack()},
		{name: "by linked list", stack: ch08.NewLinkedListStack()},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// push
			for _, x := range []int{3, 7, 5, 4} {
				if err := tt.stack.Push(x); err != nil {
					t.Fatal(err)
				}
			}
			if got, want := tt.stack.Values(), []int{3, 7, 5, 4}; !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}

			// pop
			got, err := tt.stack.Pop()
			if err != nil {
				t.Fatal(err)
			}
			if want := 4; got != want {
				t.Errorf("got: %d, want: %d", got, want)
			}
			if got, want := tt.stack.Values(), []int{3, 7, 5}; !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}

			// push
			if err := tt.stack.Push(9); err != nil {
				t.Fatal(err)
			}
			if got, want := tt.stack.Values(), []int{3, 7, 5, 9}; !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}
		})
	}
}
