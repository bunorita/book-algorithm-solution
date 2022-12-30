package ch07

import (
	"fmt"
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

// 7.3
// https://atcoder.jp/contests/agc009/tasks/agc009_a
func MultipleArray(a, b []int) int {
	n := len(a)

	var sum int
	for i := n - 1; i >= 0; i-- {
		a[i] += sum // 前回までの操作回数を足す

		var di int
		if r := a[i] % b[i]; r != 0 {
			di = b[i] - r
		}
		fmt.Printf("i=%d, d[i]=%d\n", i, di)
		sum += di
	}

	return sum
}

// ex 7.1
func CreatePairs(a, b []int) int {
	n := len(a)
	sort.Ints(a) // O(N*logN)
	sort.Ints(b)

	var i int
	for j := 0; j < n; j++ { // O(N)
		if a[i] < b[j] {
			i++
		}
	}
	return i

	// var count int
	// var nextJ int
	// for i := 0; i < n; i++ {
	// 	for j := nextJ; j < n; j++ {
	// 		if a[i] < b[j] {
	// 			count++
	// 			nextJ = j + 1
	// 			break
	// 		}
	// 	}
	// }
	// return count
}

// ex7.3
// https://atcoder.jp/contests/abc131/tasks/abc131_d
func JobsDone(jobs []*Job) bool {
	// Order BY deadline DESC
	// O(N*logN)
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Deadline < jobs[j].Deadline
	})

	t := 0 // current itme
	for _, job := range jobs {
		if t+job.Duration > job.Deadline {
			return false
		}
		t += job.Duration
	}
	return true
}

type Job struct {
	Duration int
	Deadline int
}
