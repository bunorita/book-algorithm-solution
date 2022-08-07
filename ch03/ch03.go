// full search
package ch03

import (
	"math"
)

// code 3.1
// linear serach
func Includes(a []int, v int) bool {
	for _, ai := range a {
		if ai == v {
			return true
		}
	}
	return false
}

// code 3.2
// index at which v can be found in a
func IndexOf(a []int, v int) int {
	for i, ai := range a {
		if ai == v {
			return i
		}
	}
	return -1
}

// code 3.3
// find minimum value in a
func Min(a []int) int {
	min := math.MaxInt
	for _, ai := range a {
		if ai < min {
			min = ai
		}
	}
	return min
}

// code 3.4
// find minimum sum of two numbers from each slices
// O(n^2)
func MinSum(a, b []int) int {
	min := math.MaxInt
	for _, ai := range a {
		for _, bi := range b {
			if ai+bi < min {
				min = ai + bi
			}
		}
	}
	return min
}

// PartialSumEquals returns whether partial sum of a equals w
// bit全探索
// O(2^n)
func PartialSumEquals(a []int, w int) bool {
	n := len(a)
	// n個の整数からなる集合の部分集合は2^n通りある. 2^n == (1<<n)
	for bit := 0; bit < (1 << n); bit++ { // bit=0,1,..,2^n-1
		// calculate partial sum
		var sum int
		for i, ai := range a {
			if bit&(1<<i) > 0 { // i桁目(i番目要素)が部分集合に含まれている
				sum += ai
			}
		}

		if sum == w {
			return true
		}
	}
	return false
}
