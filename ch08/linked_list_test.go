package ch08_test

import (
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch08"
)

func TestListInsert(t *testing.T) {
	l := ch08.NewList("sato", "suzuki")

	yamada := ch08.NewNode("yamada")
	sato := l.GetNodeByName("sato")
	l.Insert(yamada, sato)

	want := "sato -> yamada -> suzuki -> "
	if got := l.String(); got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func TestListString(t *testing.T) {
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

			got := ch08.NewList(tt.names...).String()
			if got != tt.want {
				t.Errorf("got: %q, want: %q\n", got, tt.want)
			}
		})
	}
}
