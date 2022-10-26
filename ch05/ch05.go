package ch05

import (
	"fmt"
	"math"
	"sync"

	"github.com/bunorita/book-algorithm-solution/ch03"
	"github.com/bunorita/book-algorithm-solution/intutil"
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

// 5.9 Optimize intervals
// JOI 2016 A Oranges
func BoxingOranges(sizes []int) int {
	n := len(sizes)
	// dp[i] 先頭からi個の要素、区間[0,i) について
	// いくつかの区間に分割した時の最小コスト（各区間コストの総和の最小値）
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	// [0,j) = [0,i) + [i,j)
	// chmin(dp[j], dp[i] + cost(i,j))
	for j := 0; j <= n; j++ {
		for i := 0; i < j; i++ {
			chmin(&dp[j], dp[i]+costOfBox(sizes[i:j]))
		}
	}

	return dp[n]
}

func costOfBox(sizes []int) int {
	k := 5 // fixed cost
	max := intutil.Max(sizes...)
	min := intutil.Min(sizes...)
	return k + (max-min)*len(sizes)
}

// ex 5.1
func Vacation(x [][3]int) int {
	n := len(x)
	dp := make([][3]int, n+1)
	// dp[i][j] 最初のi日間の行動のうち、最後の日の活動が、
	// j = 0 のときは「海での泳ぎ」、
	// j = 1 のときは「虫捕り」、
	// j = 2 のときは「宿題」

	for i := range dp {
		if i == 0 {
			continue
		}
		for j := range dp[i] { // 0,1,2
			for k := range dp[i] { // 0,1,2
				if k == j {
					continue
				}
				chmax(&dp[i][j], dp[i-1][k]+x[i-1][j])
			}
		}
		// j := 0 // i-1日目b or c, i日目a
		// chmax(&dp[i][0], dp[i-1][1]+a[i-1][0]) // i-1日目b, => i日目a
		// chmax(&dp[i][0], dp[i-1][2]+a[i-1][0]) // i-1日目c, => i日目a

		// // j := 1 // i-1日目c or a, i日目b
		// chmax(&dp[i][1], dp[i-1][2]+a[i-1][1]) // i-1日目c, => i日目b
		// chmax(&dp[i][1], dp[i-1][0]+a[i-1][1]) // i-1日目a, => i日目b

		// // j := 2 // i-1日目a or b, i日目c
		// chmax(&dp[i][2], dp[i-1][0]+a[i-1][2]) // i-1日目c, => i日目b
		// chmax(&dp[i][2], dp[i-1][1]+a[i-1][2]) // i-1日目a, => i日目b
	}

	return intutil.Max(dp[n][:]...)
}

// ex5.2
// aからいくつか選んだ数の総和がwに一致する
func PartialSumEquals(a []int, w int) bool {
	n := len(a)

	// dp[i][j] 最初のi個の整数の中からいくつか選んだ総和がjに等しくできる
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, w+1)
	}

	dp[0][0] = true // 0個選んだ総和が0
	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ {
			// a[i] を選ばない場合
			dp[i+1][j] = dp[i][j]

			// a[i] を選ぶ場合: i個から選んでjに等しくなるか <=> i-1個から選んでj-a[i]に等しくなるか
			if j-a[i] >= 0 {
				dp[i+1][j] = dp[i][j-a[i]]
			}
		}
	}

	return dp[n][w]
}

// ex5.3
// aは正の整数配列
// aからいくつか選んだ数の総和（1以上w以下）は何通りあるか
func CountPartialSumLTE(a []int, w int) int {
	// 前半はex5.2と同じ
	n := len(a)
	// dp[i][j] 最初のi個の整数の中からいくつか選んだ総和がjに等しくできる
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, w+1)
	}

	dp[0][0] = true // 0個選んだ総和が0
	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ {
			// a[i]を選ばない場合
			dp[i+1][j] = dp[i][j]

			// a[i] を選ぶ場合: i個から選んでjに等しくなるか <=> i-1個から選んでj-a[i]に等しくなるか
			if j-a[i] >= 0 {
				dp[i+1][j] = dp[i][j-a[i]]
			}
		}
	}

	// for debug
	// for i := range dp {
	// 	for j := range dp[i] {
	// 		fmt.Printf("%v からいくつか選んだ総和が%dに等しくなる? %t\n", a[:i], j, dp[i][j])
	// 	}
	// }

	var count int
	for _, dpnj := range dp[n][1:] { // dp[n][0] は除外する(j>=1)
		if dpnj {
			count++
		}
	}
	return count
}

// ex5.4
// aからk個以下の整数を選んでwに一致させることができるか
func KIntsOrLessPartialSumEquals(a []int, K int, w int) bool {
	// 前半はex5.2と同じ
	n := len(a)
	// dp[i][j][k]
	// 最初の i 個の整数の中から、総和が j がなるように選ぶことができるかどうかを表すブール値 (true / false)。
	// ただし選ぶ個数が k 個以内となるようにする
	dp := make([][][]bool, n+1)
	for i := range dp {
		dp[i] = make([][]bool, w+1)
		for j := range dp[i] {
			dp[i][j] = make([]bool, K+1)
		}
	}

	dp[0][0][0] = true // 0個選んだ総和が0

	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ {
			for k := 0; k <= K; k++ {
				// a[i] を選ばない場合
				dp[i+1][j][k] = dp[i][j][k]

				// k個以内で a[i] を選ぶ場合: i個から選んでjに等しくなるか <=> i-1個から選んでj-a[i]に等しくなるか
				if j-a[i] >= 0 && k-1 >= 0 {
					dp[i+1][j][k] = dp[i][j-a[i]][k-1]
				}
			}
		}

	}

	return dp[n][w][K]
}
