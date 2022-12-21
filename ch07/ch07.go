package ch07

import (
	"sort"
)

type Interval struct {
	Begin int
	End   int
}

// 7.2
// https://atcoder.jp/contests/typical-algorithm/tasks/typical_algorithm_b
func IntervalScheduling(intervals []*Interval) int {
	// O(N*logN)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].End < intervals[j].End
	})

	prevEnd := 0
	selected := []*Interval{}
	for _, intvl := range intervals {
		if intvl.Begin >= prevEnd {
			selected = append(selected, intvl)
			prevEnd = intvl.End
		}
	}

	return len(selected)
}
