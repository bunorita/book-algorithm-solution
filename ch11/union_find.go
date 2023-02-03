package ch11

import "sort"

type UnionFind struct {
	n   int
	par []int
	siz []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		n:   n,
		par: make([]int, n),
		siz: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.par[i] = -1
		uf.siz[i] = 1
	}
	return uf
}

func (uf *UnionFind) Root(x int) int {
	if uf.isRoot(x) {
		return x
	}
	root := uf.Root(uf.par[x])
	uf.par[x] = root // path compression 経路圧縮
	return root
}

func (uf *UnionFind) isRoot(x int) bool {
	return uf.par[x] == -1
}

func (uf *UnionFind) IsSame(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

func (uf *UnionFind) Unite(x, y int) bool {
	if uf.IsSame(x, y) {
		return false
	}

	// union by size
	// サイズの小さい方(y)のrootを、大きい方(x)のrootの子頂点にする
	x, y = uf.Root(x), uf.Root(y)
	// meet: size of x >= size of y
	if uf.siz[x] < uf.siz[y] {
		x, y = y, x // swap
	}

	uf.par[y] = x
	uf.siz[x] += uf.siz[y]
	return true
}

func (uf *UnionFind) Unions() [][]int {
	roots := make([]int, uf.n)
	for i := 0; i < uf.n; i++ {
		roots[i] = uf.Root(i)
	}

	root2union := make(map[int][]int)
	for i, root := range roots {
		root2union[root] = append(root2union[root], i)
	}

	unions := [][]int{}
	for _, union := range root2union {
		unions = append(unions, union)
	}
	sort.Slice(unions, func(i int, j int) bool {
		return unions[i][0] < unions[j][0]
	})
	return unions
}

func (uf *UnionFind) UnionCount() int {
	var count int
	for i := 0; i < uf.n; i++ {
		if uf.isRoot(i) {
			count++
		}
	}
	return count
}

// xを含む連結部分のサイズ
func (uf *UnionFind) UnionSize(x int) int {
	return uf.siz[uf.Root(x)]
}
