package ch05

import (
	"math"
	"sync"

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

// 5.4
// push-based
func Frog1PushBased(h []int) []int {
	dp := make([]int, len(h))
	for i := range dp {
		dp[i] = math.MaxInt // initialize
	}

	dp[0] = 0
	for i := 0; i < len(h)-1; i++ {
		chmin(&dp[i+1], dp[i]+AbsInt(h[i+1]-h[i])) // jump from i to i+1
		if i < len(h)-2 {
			chmin(&dp[i+2], dp[i]+AbsInt(h[i+2]-h[i])) // jump from i to i+2
		}
	}
	return dp
}

// 5.5
// recursive
func Frog1R(h []int) []int {
	mins := make([]int, len(h))
	for i := range h {
		mins[i] = frog1R(h, i)
	}
	return mins
}

// returns minimum cost to reach h[i]
func frog1R(h []int, i int) int {
	if i == 0 {
		return 0
	}
	min := math.MaxInt
	chmin(&min, frog1R(h, i-1)+AbsInt(h[i]-h[i-1])) // jump from i-1 to i
	if i > 1 {
		chmin(&min, frog1R(h, i-2)+AbsInt(h[i]-h[i-2])) // jump from i-2 to i
	}
	return min
}

// 5.6
// recursive & memoization
func Frog1RM(h []int) []int {
	mins := make([]int, len(h))
	for i := range h {
		mins[i] = frog1RM(h, i)
	}
	return mins
}

var frog1Memo = make(map[int]int)
var mu sync.RWMutex

func frog1RM(h []int, i int) int {
	if i == 0 {
		return 0
	}
	min := math.MaxInt

	// jump from i-1 to i
	mu.RLock()
	min_i_1, ok := frog1Memo[i-1]
	mu.RUnlock()
	if !ok {
		min_i_1 = frog1R(h, i-1)
		mu.Lock()
		frog1Memo[i-1] = min_i_1
		mu.Unlock()
	}
	chmin(&min, min_i_1+AbsInt(h[i]-h[i-1]))

	if i > 1 {
		// jump from i-2 to i
		mu.RLock()
		min_i_2, ok := frog1Memo[i-2]
		mu.RUnlock()
		if !ok {
			min_i_2 = frog1R(h, i-2)
			mu.Lock()
			frog1Memo[i-2] = min_i_2
			mu.Unlock()
		}
		chmin(&min, min_i_2+AbsInt(h[i]-h[i-2]))
	}
	return min
}
