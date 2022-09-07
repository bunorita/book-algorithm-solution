package ch05

import (
	"fmt"
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

type Item struct {
	Weight int
	Value  int
}

// 5.7
// N個の品物から、重さの総和がWを超えないように品物を選ぶ。
// 選んだ品物の価値の総和の最大値を返す。
func Knapsack(items []*Item, W int) int {
	N := len(items)
	weights, values := make([]int, N), make([]int, N)
	for i, item := range items {
		weights[i] = item.Weight
		values[i] = item.Value
	}

	// dp[i+1][w] i個の品物{0,1,...i-1}までの中から重さがwを超えないように選んだ時の、価値の総和の最大値
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}
	// 品物が全くない状態では重さも価値も0
	for w := 0; w <= W; w++ {
		dp[0][w] = 0
	}

	for i := 0; i < N; i++ {
		for w := 0; w <= W; w++ {
			// i番目の品物を選ぶ場合
			if weights[i] <= w {
				chmax(&dp[i+1][w], dp[i][w-weights[i]]+values[i])
			}
			// i番目の品物を選ばない場合
			chmax(&dp[i+1][w], dp[i][w])
		}
	}

	for i, dpi := range dp {
		fmt.Printf("i=%d %v\n", i, dpi)
	}
	return dp[N][W]
}

func chmax(a *int, b int) {
	if b > *a {
		*a = b
	}
}

// 5.8 edit distance
func EditDist(s, t string) int {
	// dp[i][j] sの最初のiも自分と、tの最初のj文字分との編集距離
	// edit distance between s[:i] and t[:j]
	// dp: len(s)+1 * len(t)+1

	// initialize
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = 0 // edit distance between "" and "" is zero

	for i := 0; i <= len(s); i++ {
		for j := 0; j <= len(t); j++ {
			// 変更操作: sのi文字目とtのj文字目を対応させる
			if i > 0 && j > 0 {
				if s[i-1] == t[j-1] {
					chmin(&dp[i][j], dp[i-1][j-1])
				} else {
					chmin(&dp[i][j], dp[i-1][j-1]+1)
				}
			}
			// 削除操作：sのi文字目を削除
			if i > 0 {
				chmin(&dp[i][j], dp[i-1][j]+1)
			}
			// 挿入操作：=> tのj文字目を削除
			if j > 0 {
				chmin(&dp[i][j], dp[i][j-1]+1)
			}
		}
	}

	// for _, row := range dp {
	// 	fmt.Println(row)
	// }

	return dp[len(s)][len(t)]
}
