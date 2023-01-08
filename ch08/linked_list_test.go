package ch08_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch08"
)

func TestLinkedListString(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name  string
		names []string
		want  string
	}{
		{
			name:  "a name",
			names: []string{"sato"},
			want:  "sato -> ",
		},
		{
			name: "6 names",
			names: []string{
				"sato",
				"suzuki",
				"takahashi",
				"ito",
				"watanabe",
				"yamamoto",
			},
			want: "sato -> suzuki -> takahashi -> ito -> watanabe -> yamamoto -> ",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch08.NewLinkedList(tt.names...).String()
			if got != tt.want {
				t.Errorf("got: %q, want: %q\n", got, tt.want)
			}
		})
	}
}

func TestLinkedListErase(t *testing.T) {
	t.Parallel()

	lst := ch08.NewLinkedList("sato", "yamada", "suzuki")
	yamada := lst.GetNodeByValue("yamada")
	lst.Erase(yamada)

	want := "sato -> suzuki -> "
	if got := lst.String(); got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func TestLinkedListAppend(t *testing.T) {
	t.Parallel()

	lst := ch08.NewLinkedList("sato", "suzuki")
	yamada := ch08.NewLinkedListNode("", "yamada")
	lst.Append(yamada)

	want := "sato -> suzuki -> yamada -> "
	if got := lst.String(); got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func TestLinkedListGetNode(t *testing.T) {
	t.Parallel()

	lst := ch08.NewLinkedList("sato", "yamada", "suzuki")
	sato := lst.GetNodeByValue("sato")
	yamada := lst.GetNodeByValue("yamada")
	suzuki := lst.GetNodeByValue("suzuki")

	tests := []*struct {
		i    int
		want *ch08.LinkedNode
	}{
		{0, sato},
		{1, yamada},
		{2, suzuki},
		{3, nil},
		{-1, nil},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("i=%d", tt.i), func(t *testing.T) {
			t.Parallel()

			got := lst.GetNode(tt.i)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v\n", got, tt.want)
			}
		})
	}
}
