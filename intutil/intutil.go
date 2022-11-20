package intutil

import "math"

func Max(x ...int) int {
	max := math.MinInt
	for _, xi := range x {
		if xi > max {
			max = xi
		}
	}
	return max
}

func Min(x ...int) int {
	min := math.MaxInt
	for _, xi := range x {
		if xi < min {
			min = xi
		}
	}
	return min
}

func Chmin(a *int, b int) {
	if b < *a {
		*a = b
	}
}

func Chmax(a *int, b int) {
	if b > *a {
		*a = b
	}
}
