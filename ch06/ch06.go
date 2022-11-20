package ch06

import (
	"math"
	"sort"

	"github.com/bunorita/book-algorithm-solution/intutil"
)

// 6.1
// a is sorted
func BinarySearch(a []int, key int) int {
	l, r := 0, len(a)-1

	for l <= r {
		mid := l + (r-l)/2
		if key == a[mid] {
			return mid
		} else if key < a[mid] {
			r = mid - 1
		} else { // a[mid] < key
			l = mid + 1
		}
	}

	return -1
}

// 6.2
// generalized
// returns minimum i that satisfies a[i] >= key
func BinarySearchGen(a []int, key int) int {
	// p(left) => false, p(right) => true になるように
	p := func(m int) bool {
		if m < 0 {
			return false
		}
		if m >= len(a) {
			return true
		}
		return a[m] >= key
	}
	left, right := -1, len(a)

	for right-left > 1 {
		mid := left + (right-left)/2
		if p(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

// 6.4
// find minimum sum(>=k) of two numbers from each slices
// O(NlogN)
// N = len(a) = len(b)
func MinSum(a, b []int, k int) int {
	// O(NlogN)
	sort.Ints(a)
	sort.Ints(b)
	bmax := b[len(b)-1]

	min := math.MaxInt
	for i := range a { // O(N)
		if bmax < k-a[i] {
			continue
		}
		// bmax >= k-ai
		//   => b[j] exists: b[j] >= k-a[i]
		j := BinarySearchGen(b, k-a[i]) // O(logN)
		intutil.Chmin(&min, a[i]+b[j])
	}
	return min
}

// 6.5 射撃王
// h: 初期高度
// s: 毎秒の増加高度
// ペナルティ(全ての風船を割った時の最大高度)の最小値を求める
func Shooting(h, s []int) int {
	// 全ての風船を割ることができるペナルティxの最小値を求める

	n := len(h) // n = len(h) = len(s)
	left := 0
	var right int // 二分探索の上限値(全ての風船を割れるペナルティ)
	for i := 0; i < n; i++ {
		intutil.Chmax(&right, h[i]+s[i]*n)
	}

	// 判定する: ペナルティx以下で全ての風船を割ることができるか?
	p := func(x int) bool {
		t := make([]float64, n) // 各風船を割らなければいけない残り時間
		for i := 0; i < n; i++ {
			// そもそもペナルティが初期高度よりも低かったらfalse
			if x < h[i] {
				return false
			}
			t[i] = float64(x-h[i]) / float64(s[i]) // 残り上昇距離 / 毎秒上昇距離
		}

		sort.Float64s(t) // 時間制限が差し迫っている順にソートする
		for i := range t {
			if t[i] < float64(i) { // 時間切れ発生の場合はfalse
				return false
			}
		}
		return true
	}

	for right-left > 1 {
		mid := (left + right) / 2 // = left+(right-left)/2
		if p(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
