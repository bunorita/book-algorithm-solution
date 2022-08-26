package ch05

import (
	"math"

	"github.com/bunorita/book-algorithm-solution/ch03"
)

// code 5.1
func Frog1(h []int) []int {
	dp := make([]int, len(h))
	for i := range dp {
		dp[i] = math.MaxInt // initialize
	}

	dp[0] = 0
	for i := 1; i < len(h); i++ {
		if i == 1 {
			dp[1] = AbsInt(h[1] - h[0])
			continue
		}
		// i >= 2
		dp[i] = ch03.Less(
			dp[i-1]+AbsInt(h[i]-h[i-1]), // jump from i-1 to i
			dp[i-2]+AbsInt(h[i]-h[i-2]), // jump from i-2 to i
		)
	}
	return dp
}

func AbsInt(x int) int {
	return int(math.Abs(float64(x)))
}

// 5.2
// choose minimum
func chmin(a *int, b int) {
	if b < *a {
		*a = b
	}
}

// 5.3
func Frog1Relaxation(h []int) []int {
	dp := make([]int, len(h))
	for i := range dp {
		dp[i] = math.MaxInt // initialize
	}

	dp[0] = 0
	for i := 1; i < len(h); i++ {
		chmin(&dp[i], dp[i-1]+AbsInt(h[i]-h[i-1])) // jump from i-1 to i
		if i > 1 {
			chmin(&dp[i], dp[i-2]+AbsInt(h[i]-h[i-2])) // jump from i-2 to i
		}
	}
	return dp
}
