package ch11

import "strconv"

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
		edges[i][0]--
		edges[i][1]--
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

// ex 11.3
// https://atcoder.jp/contests/abc049/tasks/arc065_b
func TransportNetwork(n int, roads, rails [][2]int) []int {
	// k, l := len(roads), len(rails)

	// 1-based index => 0-based index
	for i := range roads {
		roads[i][0]--
		roads[i][1]--
	}
	for i := range rails {
		rails[i][0]--
		rails[i][1]--
	}

	ufRoad, ufRail := NewUnionFind(n), NewUnionFind(n)
	for _, road := range roads {
		ufRoad.Unite(road[0], road[1])
	}
	for _, rail := range rails {
		ufRail.Unite(rail[0], rail[1])
	}

	m := map[string]int{}
	key := func(roadRt, railRt int) string { return strconv.Itoa(roadRt) + "-" + strconv.Itoa(railRt) }
	for i := 0; i < n; i++ {
		m[key(ufRoad.Root(i), ufRail.Root(i))]++
	}
	counts := make([]int, n)
	for i := 0; i < n; i++ {
		counts[i] = m[key(ufRoad.Root(i), ufRail.Root(i))]
	}

	return counts
}
