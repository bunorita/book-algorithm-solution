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

// ex 8.7
// O(N)
func SumEquals(a, b []int, k int) bool {
	m := make(map[int]bool) // whether the number exists in a
	for i := range a {      // O(N)
		m[a[i]] = true
	}

	// ai + bj = k <=> ai = k - bj
	for j := range b { // O(N)
		if m[k-b[j]] {
			return true
		}
	}
	return false
}
