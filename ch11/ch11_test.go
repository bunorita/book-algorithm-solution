package ch11_test

import (
	"reflect"
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch11"
)

func TestBridge(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name  string
		n     int
		edges [][2]int
		want  int
	}{
		{
			name: "case1",
			n:    7,
			edges: [][2]int{
				{0, 2},
				{1, 6},
				{2, 3},
				{3, 4},
				{3, 5},
				{4, 5},
				{5, 6},
			},
			want: 4,
		},
		{
			name: "case2",
			n:    3,
			edges: [][2]int{
				{0, 1},
				{0, 2},
				{1, 2},
			},
			want: 0,
		},
		{
			name: "case3",
			n:    6,
			edges: [][2]int{
				{0, 1},
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch11.Bridge(tt.n, tt.edges)
			if got != tt.want {
				t.Errorf("got=%d, want=%d", got, tt.want)
			}
		})
	}
}

func TestDecayedBridge(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name  string
		n     int
		edges [][2]int
		want  []int
	}{
		{
			name: "case1",
			n:    4,
			edges: [][2]int{
				{1, 2},
				{3, 4},
				{1, 3},
				{2, 3},
				{1, 4},
			},
			want: []int{0, 0, 4, 5, 6},
		},
		{
			name: "case2",
			n:    6,
			edges: [][2]int{
				{2, 3},
				{1, 2},
				{5, 6},
				{3, 4},
				{4, 5},
			},
			want: []int{8, 9, 12, 14, 15},
		},
		{
			name: "case3",
			n:    2,
			edges: [][2]int{
				{1, 2},
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch11.DecayedBridge(tt.n, tt.edges)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
