// SumUpTo
package ch04

import (
	"fmt"
	"sync"
)

// code 4.2
func SumUpTo(n int) int {
	fmt.Printf("SumUpTo(%d) called\n", n)
	if n == 0 {
		return 0
	}
	result := n + SumUpTo(n-1)
	fmt.Printf("sum of integers up to %d: %d\n", n, result)
	return result
}

// code 4.4
func GCD(m, n int) int {
	if n == 0 {
		return m
	}
	r := m % n
	return GCD(n, r)
}

// code 4.5
// Fib returns Fibonacci number
// recursive ver
func Fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// code 4.7
// loop ver
// O(n)
func FibLoop(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	fn_1, fn_2 := 1, 0 // f(n-1), f(n-2)
	var fn int
	for i := 2; i <= n; i++ {
		fn = fn_1 + fn_2
		fn_1, fn_2 = fn, fn_1
	}
	return fn
}

var mu sync.RWMutex // prevent concurrent map read and map write
var fib = map[int]int{
	0: 0,
	1: 1,
}

// code 4.8
// recursive with memoization
func FibMemo(n int) int {
	// read from memo
	mu.RLock()
	fn, ok := fib[n]
	mu.RUnlock()

	if !ok {
		fn = FibMemo(n-1) + FibMemo(n-2)
		// write to memo
		mu.Lock()
		fib[n] = fn
		mu.Unlock()
	}
	return fn
}
