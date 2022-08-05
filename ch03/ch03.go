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
	idx := -1
	for i, ai := range a {
		if ai == v {
			idx = i
			break
		}
	}
	return idx
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
