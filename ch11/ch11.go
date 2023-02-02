package ch11

// ex 11.1
// https://atcoder.jp/contests/abc075/tasks/abc075_c
func Bridge(n int, edges [][2]int) int {
	var count int

	for i := range edges {
		// i番目の辺を除外して、連結であるか調べる
		uf := NewUnionFind(n)
		for _, edge := range edges[:i] {
			uf.Unite(edge[0], edge[1])
		}
		for _, edge := range edges[i+1:] {
			uf.Unite(edge[0], edge[1])
		}
		if uf.UnionCount() > 1 {
			count++
		}
	}
	return count
}
