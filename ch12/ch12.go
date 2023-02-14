package ch12

import (
	"log"

	"github.com/bunorita/book-algorithm-solution/ch09"
)

func InsertionSort(ap *[]int) {
	a := *ap
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
}

func MergeSort(a *[]int) {
	mergeSort(a, 0, len(*a))
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

func QuickSort(ap *[]int) {
	quickSort(ap, 0, len(*ap))
}

func quickSort(ap *[]int, left, right int) {
	if right-left <= 1 {
		return
	}

	a := *ap

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

	quickSort(ap, left, i)
	quickSort(ap, i+1, right)
}

func HeapSort(ap *[]int) {
	a := *ap
	n := len(a)

	h := ch09.NewHeap()
	h.Push(a...)

	for i := n - 1; i >= 0; i-- { // 昇順で返すので後ろから詰める
		x, err := h.Pop()
		if err != nil {
			log.Fatal(err)
		}
		a[i] = x
	}
}

// 配列それ自体をヒープにする in-place
func HeapSort2(p *[]int) {
	a := *p
	n := len(a)

	// aをヒープにする
	// max parent index = (n-1)/2
	for i := (n - 1) / 2; i >= 0; i-- {
		heapify(p, i, n) // heapify under i
	}

	for i := n - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0] // ヒープの最大値 a[0] を右詰めしていく
		heapify(p, 0, i)        // ヒープサイズはiに
	}
}

// k番目の頂点を根とする部分木について、ヒープ条件を満たすようにする
// NOTE: a[0:n]部分についてのみ考える。 n != len(a)
func heapify(p *[]int, k, n int) {
	a := *p

	c1 := 2*k + 1 // left child
	if c1 >= n {  // if child1 does not exist
		return
	}

	c := c1                      // child to swap
	c2 := 2*k + 2                // right child
	if c2 < n && a[c1] < a[c2] { // if c2 exists & c2 > c1
		c = c2
	}

	if a[k] >= a[c] { // if parent > child
		return
	}

	a[k], a[c] = a[c], a[k]
	heapify(p, c, n)
}

func BucketSort(p *[]int) {
	const max = 100_000
	a := *p
	n := len(a)

	count := make([]int, max) // count[v] 値vの個数
	for i := range a {
		count[a[i]]++
	}

	countLTE := make([]int, max) // countLTE[v] v以下の値の個数
	for v := 0; v < max; v++ {
		if v == 0 {
			countLTE[v] = count[v]
			continue
		}
		countLTE[v] = countLTE[v-1] + count[v]
	}

	sorted := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		j := countLTE[a[i]] - 1
		sorted[j] = a[i]
		countLTE[a[i]]--
	}
	*p = sorted
}
