package ch03

func LinearSearch(a []int, v int) bool {
	for _, ai := range a {
		if ai == v {
			return true
		}
	}
	return false
}
