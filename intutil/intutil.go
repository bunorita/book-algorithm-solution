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

// prevent overflows
func SafeAdd(a, b int) int {
	if a >= 0 && b >= 0 &&
		math.MaxInt-a < b { // overflow: MaxInt < a+b
		return math.MaxInt
	}
	if a < 0 && b < 0 &&
		a < math.MinInt-b { // overflow: a+b < MinInt
		return math.MinInt
	}
	return a + b
}
