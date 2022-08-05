package ch03

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
		}
	}
	return idx
}
