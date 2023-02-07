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
	al := make([]int, mid-left)
	copy(al, (*a)[left:mid])

	mergeSort(a, mid, right)
	ar := make([]int, right-mid)
	copy(ar, (*a)[mid:right])

	li, ri := 0, 0
	for i := left; i < right; i++ {
		if li == len(al) { // 左が空
			(*a)[i] = ar[ri]
			ri++
			continue
		} else if ri == len(ar) { // 右が空
			(*a)[i] = al[li]
			li++
			continue
		}

		// 左右とも値がある
		if al[li] <= ar[ri] { // 左の方が小さい
			(*a)[i] = al[li]
			li++
		} else { // 右の方が小さい
			(*a)[i] = ar[ri]
			ri++
		}
	}
}
