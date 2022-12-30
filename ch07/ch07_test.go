package ch07_test

import (
	"testing"

	"github.com/bunorita/book-algorithm-solution/ch07"
)

func TestIntervalScheduling(t *testing.T) {
	t.Parallel()

	// N=3
	intervals := []*ch07.Interval{
		{1, 5},
		{4, 6},
		{6, 8},
	}
	want := 2
	if got := ch07.IntervalScheduling(intervals); got != want {
		t.Errorf("got=%d, want=%d\n", got, want)
	}

}

func TestMultipleArray(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name string
		a    []int
		b    []int
		want int
	}{
		{
			name: "N=3",
			a:    []int{3, 2, 9},
			b:    []int{5, 7, 4},
			want: 7,
		},
		{
			name: "N=7",
			a:    []int{3, 4, 5, 2, 5, 5, 9},
			b:    []int{1, 1, 9, 6, 3, 8, 7},
			want: 22,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := ch07.MultipleArray(tt.a, tt.b); got != tt.want {
				t.Errorf("got=%d, want=%d\n", got, tt.want)
			}
		})
	}
}

func TestCreatePairs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a, b []int
		want int
	}{
		{
			name: "N=1. 1",
			a:    []int{1},
			b:    []int{2},
			want: 1, // (1,2)
		},
		{
			name: "N=1. 2",
			a:    []int{2},
			b:    []int{1},
			want: 0,
		},
		{
			name: "N=1. 3",
			a:    []int{1},
			b:    []int{1},
			want: 0,
		},
		{
			name: "N=2. 1",
			a:    []int{1, 2},
			b:    []int{2, 3},
			want: 2, // (1,2), (2,3)
		},
		{
			name: "N=2. 2",
			a:    []int{2, 2},
			b:    []int{1, 3},
			want: 1, //  (2,3)
		},
		{
			name: "N=3",
			a:    []int{1, 3, 5},
			b:    []int{2, 4, 6},
			want: 3, // example: (1,2), (3,4), (5,6)
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := ch07.CreatePairs(tt.a, tt.b); got != tt.want {
				t.Errorf("got=%d, want=%d\n", got, tt.want)
			}

		})
	}
}

func TestJobsDone(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		name string
		jobs []*ch07.Job
		want bool
	}{
		{
			name: "example1 N=5",
			jobs: []*ch07.Job{
				{2, 4},
				{1, 9},
				{1, 8},
				{4, 9},
				{3, 12},
			},
			want: true,
		},
		{
			name: "example2 N=3",
			jobs: []*ch07.Job{
				{334, 1000},
				{334, 1000},
				{334, 1000},
			},
			want: false,
		},
		{
			name: "example3 N=30",
			jobs: []*ch07.Job{
				{384, 8895},
				{1725, 9791},
				{170, 1024},
				{4, 11105},
				{2, 6},
				{578, 1815},
				{702, 3352},
				{143, 5141},
				{1420, 6980},
				{24, 1602},
				{849, 999},
				{76, 7586},
				{85, 5570},
				{444, 4991},
				{719, 11090},
				{470, 10708},
				{1137, 4547},
				{455, 9003},
				{110, 9901},
				{15, 8578},
				{368, 3692},
				{104, 1286},
				{3, 4},
				{366, 12143},
				{7, 6649},
				{610, 2374},
				{152, 7324},
				{4, 7042},
				{292, 11386},
				{334, 5720},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := ch07.JobsDone(tt.jobs); got != tt.want {
				t.Errorf("got: %t, want: %t\n", got, tt.want)
			}
		})

	}
}
