package ch06

// a is sorted
func BinarySearch(a []int, key int) int {
	l, r := 0, len(a)-1

	for l <= r {
		mid := l + (r-l)/2
		if key == a[mid] {
			return mid
		} else if key < a[mid] {
			r = mid - 1
		} else { // a[mid] < key
			l = mid + 1
		}
	}

	return -1
}
