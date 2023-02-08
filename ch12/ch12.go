package ch12

func InsertionSort(a []int) []int {
	n := len(a)
	for i := 1; i < n; i++ {
		v := a[i]

		// vを挿入するインデックスjを探す
		j := i
		for ; j > 0; j-- {
			if a[j-1] <= v {
				break
			}
			// a[j-1] > v
			a[j] = a[j-1] // 一つ後ろにずらす
		}
		a[j] = v
	}
	return a
}

func MergeSort(a []int) []int {
	mergeSort(&a, 0, len(a))
	return a
}

func mergeSort(a *[]int, left, right int) {
	if right-left == 1 {
		return
	}

	mid := (right + left) / 2

	mergeSort(a, left, mid)
	mergeSort(a, mid, right)

	buf := make([]int, 0)
	for i := left; i < mid; i++ {
		buf = append(buf, (*a)[i])
	}
	for i := right - 1; i >= mid; i-- { // 右側は逆順で
		buf = append(buf, (*a)[i])
	}

	li := 0
	ri := len(buf) - 1
	for i := left; i < right; i++ {
		if buf[li] <= buf[ri] {
			(*a)[i] = buf[li]
			li++
		} else {
			(*a)[i] = buf[ri]
			ri--
		}
	}
}

func QuickSort(a []int) []int {
	quickSort(&a, 0, len(a))
	return a
}

func quickSort(arr *[]int, left, right int) {
	if right-left <= 1 {
		return
	}

	a := *arr

	pi := (left + right) / 2 // pivot index
	pivot := a[pi]

	// swap pivot with right edge
	a[pi], a[right-1] = a[right-1], a[pi]

	i := left // 左詰めされたpivot未満要素の右端+1（pivot以上要素の左端）
	for j := left; j < right-1; j++ {
		if a[j] < pivot { // pivot未満のものがあったら左に詰めていく
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	// swap pivot with i
	a[i], a[right-1] = a[right-1], a[i]

	quickSort(arr, left, i)
	quickSort(arr, i+1, right)
}
