package ch08

import "strconv"

// ex 8.5
// O(N+M)
func CountCommonNumbers(a, b []int) int {
	n, m := len(a), len(b)
	ht := NewHashTable(n + m)

	for i := range a { // O(N)
		ht.Set(strconv.Itoa(a[i]), "true")
	}

	count := 0

	for i := range b { // O(M)
		if val := ht.Get(strconv.Itoa(b[i])); val == "true" {
			count++
		}
	}
	return count
}
