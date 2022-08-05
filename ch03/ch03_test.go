package ch03_test

import (
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch03"
)

func TestIncludes(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		a    []int
		v    int
		want bool
	}{
		{
			name: "includes",
			a:    []int{4, 3, 12, 7, 11},
			v:    7,
			want: true,
		},
		{
			name: "not include",
			a:    []int{4, 3, 12, 7, 11},
			v:    9,
			want: false,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch03.Includes(tt.a, tt.v)
			if got != tt.want {
				t.Errorf("want %t, but got %t\n", tt.want, got)
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		a    []int
		v    int
		want int
	}{
		{
			name: "found",
			a:    []int{4, 3, 12, 7, 11},
			v:    7,
			want: 3,
		},
		{
			name: "not found",
			a:    []int{4, 3, 12, 7, 11},
			v:    9,
			want: -1,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch03.IndexOf(tt.a, tt.v)
			if got != tt.want {
				t.Errorf("want %d, but got %d\n", tt.want, got)
			}
		})
	}
}

func TestMin(t *testing.T) {
	t.Parallel()

	a := []int{4, 3, 12, 7, 11}
	want := 3
	got := ch03.Min(a)
	if got != want {
		t.Errorf("want %d, but got %d\n", want, got)
	}
}

func TestMinSum(t *testing.T) {
	t.Parallel()

	a := []int{4, 3, 12, 7, 11}
	b := []int{8, 1, 5, 2, 9}
	want := 4
	got := ch03.MinSum(a, b)
	if got != want {
		t.Errorf("want %d, but got %d\n", want, got)
	}
}

func TestPartialSumEquals(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		a    []int
		w    int
		want bool
	}{
		{
			name: "true",
			a:    []int{1, 2, 4, 5, 11},
			w:    10,
			want: true,
		},
		{
			name: "false",
			a:    []int{1, 5, 8, 11},
			w:    10,
			want: false,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch03.PartialSumEquals(tt.a, tt.w)
			if got != tt.want {
				t.Errorf("want %t, but got %t\n", tt.want, got)
			}
		})
	}
}
