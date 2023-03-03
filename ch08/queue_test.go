package ch08_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch08"
)

func TestQueue(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name string
		q    ch08.IQueue[int]
	}{
		{name: "by array", q: ch08.NewQueue[int]()},
		{name: "by linked list", q: ch08.NewLinkedListQueue()},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// enqueue
			for _, x := range []int{3, 5, 7} {
				if err := tt.q.Enqueue(x); err != nil {
					t.Fatal(err)
				}
			}
			if got, want := tt.q.Values(), []int{3, 5, 7}; !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}

			// dequeue
			got, err := tt.q.Dequeue()
			if err != nil {
				t.Fatal(err)
			}
			if want := 3; got != want {
				t.Errorf("got: %d, want: %d", got, want)
			}
			if got, want := tt.q.Values(), []int{5, 7}; !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}

			// enquue
			tt.q.Enqueue(9)
			if got, want := tt.q.Values(), []int{5, 7, 9}; !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}

		})
	}

}
