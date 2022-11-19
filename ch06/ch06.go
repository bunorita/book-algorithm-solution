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
