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

func TestLargestIndexOf(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		a    []int
		v    int
		want int
	}{
		{
			name: "found once",
			a:    []int{4, 3, 12, 7, 11},
			v:    7,
			want: 3,
		},
		{
			name: "found twice",
			a:    []int{4, 3, 12, 7, 11, 7, 9},
			v:    7,
			want: 5,
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

			got := ch03.LargestIndexOf(tt.a, tt.v)
			if got != tt.want {
				t.Errorf("want %d, but got %d\n", tt.want, got)
			}
		})
	}
}

func TestCount(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		a    []int
		v    int
		want int
	}{
		{
			name: "found once",
			a:    []int{4, 3, 12, 7, 11},
			v:    7,
			want: 1,
		},
		{
			name: "found twice",
			a:    []int{4, 3, 12, 7, 11, 7, 9},
			v:    7,
			want: 2,
		},
		{
			name: "not found",
			a:    []int{4, 3, 12, 7, 11},
			v:    9,
			want: 0,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch03.Count(tt.a, tt.v)
			if got != tt.want {
				t.Errorf("want %d, but got %d\n", tt.want, got)
			}
		})
	}
}

func TestSecondSmallest(t *testing.T) {
	t.Parallel()

	a := []int{1, 4, 3, 12, 7, 11}
	want := 3
	got := ch03.SecondSmallest(a)
	if got != want {
		t.Errorf("want %d, but got %d\n", want, got)
	}
}

func TestLargestDiff(t *testing.T) {
	t.Parallel()

	a := []int{4, 3, 12, 7, 11}
	want := 9
	got := ch03.LargestDiff(a)
	if got != want {
		t.Errorf("want %d, but got %d\n", want, got)
	}
}

func TestDividableCount(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		a    []int
		want int
	}{
		{
			name: "cannot divide",
			a:    []int{1, 2},
			want: 0,
		},
		{
			name: "can divide once",
			a:    []int{2, 4},
			want: 1,
		},
		{
			name: "can divide twice",
			a:    []int{4, 8},
			want: 2,
		},
		{
			name: "can divide thrice",
			a:    []int{8, 16},
			want: 3,
		},
		{
			name: "example",
			a:    []int{8, 12, 40},
			want: 2,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch03.DividableCount(tt.a)
			if got != tt.want {
				t.Errorf("want %d, but got %d\n", tt.want, got)
			}
		})
	}
}

func TestCountOfTheComb(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name       string
		k, n, want int
	}{
		{
			name: "case1",
			k:    2,
			n:    2,
			want: 6,
		},
		{
			name: "case2",
			k:    5,
			n:    15,
			want: 1,
		},
		{
			name: "case3",
			k:    10,
			n:    17,
			want: 87,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ch03.CountOfTheComb(tt.k, tt.n)
			if got != tt.want {
				t.Errorf("want %d, but got %d\n", tt.want, got)
			}
		})
	}
}

func TestSumOfTheEveryAddition(t *testing.T) {
	t.Parallel()

	s := "125"
	want := 176
	got := ch03.SumOfTheEveryAddition(s)
	if got != want {
		t.Errorf("want %d, but got %d\n", want, got)
	}
}

func TestEvalAddition(t *testing.T) {
	t.Parallel()

	tests := []struct {
		expr string
		want int
	}{
		{
			expr: "125",
			want: 125,
		},
		{
			expr: "1+25",
			want: 26,
		},
		{
			expr: "12+5",
			want: 17,
		},
		{
			expr: "1+2+5",
			want: 8,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.expr, func(t *testing.T) {
			t.Parallel()

			got := ch03.EvalAddition(tt.expr)
			if got != tt.want {
				t.Errorf("want %d, but got %d", tt.want, got)
			}
		})
	}
}
