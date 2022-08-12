// SumUpTo
package ch04

import "fmt"

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
