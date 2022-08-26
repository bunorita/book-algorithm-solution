package ch05

import (
	"math"

	"github.com/bunorita/book-algorithm-solution/ch03"
)

// code 5.1
func Frog1(h []int) []int {
	dp := make([]int, len(h))
	for i := range dp {
		dp[i] = math.MaxInt
	}

	for i := range h {
		if i == 0 {
			dp[0] = 0
			continue
		}
		if i == 1 {
			dp[1] = h[1] - h[0]
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
