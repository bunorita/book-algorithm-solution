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

// ex 11.2
// https://atcoder.jp/contests/abc120/tasks/abc120_d
func DecayedBridge(n int, edges [][2]int) []int {
	m := len(edges)

	// テストデータ 1-based index => 0-based へ変換
	for i := range edges {
		edges[i][0] -= 1
		edges[i][1] -= 1
	}

	uf := NewUnionFind(n)

	// 逆順に計算する (全ての橋が落ちた => 全ての橋が残っている)

	// inconvenience 不便さ：橋がなくて行き来できない2島の組み合わせの数
	inconvs := make([]int, m)
	inconv := n * (n - 1) / 2 // 初期値 = combination C(n,2)
	// i = m-1, ..., 1, 0 の順で橋を架けていく
	for i := m - 1; i >= 0; i-- {
		inconvs[i] = inconv

		a, b := edges[i][0], edges[i][1]

		if uf.IsSame(a, b) { // a, bは既に行き来可能なので、直接橋を架けても不便さ変わらず
			continue
		}

		conv := uf.UnionSize(a) * uf.UnionSize(b) // a,bに橋を架けると行き来できるようになる2島の組み合わせの数
		inconv -= conv
		uf.Unite(a, b)
	}

	return inconvs
}
