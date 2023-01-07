package ch08_test

import (
	"fmt"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch08"
)

func TestHashTableAccessors(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		key, want string
	}{
		{"red", "apple"},
		{"blue", "sky"},
		{"green", "leaf"},
		{"yellow", "lemon"},
		{"purple", "rain"},
	}
	h := ch08.NewHashTable(3)
	for _, tt := range tests {
		h.Set(tt.key, tt.want)
	}
	h.PrintArray()

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("key=%q", tt.key), func(t *testing.T) {
			t.Parallel()

			if got := h.Get(tt.key); got != tt.want {
				t.Errorf("got=%q, want=%q\n", got, tt.want)
			}
		})
	}
}

func TestHashTableGetEmpty(t *testing.T) {
	t.Parallel()

	h := ch08.NewHashTable(3)
	want := ""
	if got := h.Get("red"); got != want {
		t.Errorf("got=%q, want=%q\n", got, want)
	}
}

func TestHashTableUnset(t *testing.T) {
	t.Parallel()

	h := ch08.NewHashTable(1)
	h.Set("red", "sun")
	h.Set("white", "horse")
	h.Set("black", "jacket")

	h.PrintArray()

	h.Unset("white")
	want := ""
	if got := h.Get("white"); got != want {
		t.Errorf("got=%q, want=%q\n", got, want)
	}

	h.PrintArray()
}
