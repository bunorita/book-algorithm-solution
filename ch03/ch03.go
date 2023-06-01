// full search
package ch03

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/bunorita/book-algorithm-solution/intutil"
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

// code 3.6
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

// ex 3.1
func LargestIndexOf(a []int, v int) int {
	foundIdx := -1
	for i, ai := range a {
		if ai == v {
			foundIdx = i
		}
	}
	return foundIdx
}

// ex 3.2
// Count returns how many times v can be found in a
func Count(a []int, v int) int {
	var count int
	for _, ai := range a {
		if ai == v {
			count++
		}
	}
	return count
}

// ex 3.3
//
// SecondSmallest returns the second smallest value
func SecondSmallest(a []int) int {
	first, second := math.MaxInt, math.MaxInt
	for _, ai := range a {
		if ai < first {
			first, second = ai, first
		} else if ai < second {
			second = ai
		}
	}
	return second
}

// ex 3.4
// LargestDiff returns the largest difference of two numbers in a
// O(n)
func LargestDiff(a []int) int {
	max, min := math.MinInt, math.MaxInt
	for _, ai := range a {
		if ai > max {
			max = ai
		}
		if ai < min {
			min = ai
		}
	}
	return max - min
}

// ex 3.5
// DividableCount returns how many times you can divide all numbers in a by 2
func DividableCount(a []int) int {
	min := math.MaxInt
	for _, ai := range a {
		intutil.Chmin(&min, dividableCount(ai))
	}
	return min
}

// how many times n can be divided by 2
func dividableCount(n int) int {
	var exp int
	for exp = 0; n%2 == 0; exp++ {
		n /= 2
	}
	return exp
}

// less one
func Less(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// ex 3.6
// how many combinations of three integers of which each number is less than or equal to k, and the sum equals n
// https://atcoder.jp/contests/abc051/tasks/abc051_b
func CountOfTheComb(k, n int) int {
	var count int
	for x := 0; x <= k; x++ {
		for y := 0; y <= k; y++ {

			// O(n^3)
			// for z := 0; z <= k; z++ {
			// 	if x+y+z == n {
			// 		count++
			// 	}
			// }

			z := n - x - y        // want z: x+y+z = n <=> z = n-x-y
			if z >= 0 && z <= k { // if the z exists
				count++
			}
		}
	}
	return count
}

// ex 3.7
func SumOfTheEveryAddition(s string) int {
	// s = "125", n = len(s) = 3
	// 125, 1+25, 12+5, 1+2+5 : 4=2^(n-1)通り
	n := len(s)
	splitted := strings.Split(s, "")
	var sum int

	baseDigit := 1 << (n - 2) // 2^(n-2) = 0b10

	for bit := 0; bit < (1 << (n - 1)); bit++ { // 2^(n-1)通り
		var expr string
		for i, d := range splitted {
			expr += d
			currentDigit := baseDigit >> i // +を入れるか調べる場所
			if bit&(currentDigit) > 0 {
				expr += "+"
			}
		}
		fmt.Printf("bit: %b, expr: %s\n", bit, expr)
		sum += EvalAddition(expr)
	}

	return sum
}

func EvalAddition(expr string) int {
	nums := strings.Split(expr, "+")
	var sum int
	for _, num := range nums {
		i, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalln(err)
		}
		sum += i
	}
	return sum
}
