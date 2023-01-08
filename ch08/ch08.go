package ch08

import "strconv"

// ex 8.5
// O(N+M)
// a, b: has no duplicate values
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

// ex 8.6
// O(N+M)
func CountCommonNumberPairs(a, b []int) int {
	m := make(map[int]int) // number => count
	for i := range a {     // O(N)
		m[a[i]] += 1
	}

	count := 0

	for j := range b { // O(M)
		count += m[b[j]]
	}
	return count
}
