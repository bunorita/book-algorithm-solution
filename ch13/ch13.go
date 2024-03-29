package ch13

import (
	"strings"

	"github.com/bunorita/book-algorithm-solution/ch08"
	"github.com/bunorita/book-algorithm-solution/ch09"
)

func BFS(g *ch09.Graph, s int) ([]bool, error) {
	return search(g, s, ch08.NewQueue[int]())
}

func DFS(g *ch09.Graph, s int) ([]bool, error) {
	return search(g, s, ch08.NewStack())
}

type set interface {
	Push(x int) error
	Pop() (int, error)
	IsEmpty() bool
}

func search(g *ch09.Graph, s int, todo set) ([]bool, error) {
	seen := make([]bool, g.Size())

	// initialize
	seen[s] = true
	if err := todo.Push(s); err != nil {
		return nil, err
	}

	for !todo.IsEmpty() {
		v, err := todo.Pop()
		if err != nil {
			return nil, err
		}
		// vから辿れる頂点を全て調べる
		for _, x := range g.VerticesConnectedWith(v) {
			if seen[x] {
				continue
			}

			seen[x] = true
			if err := todo.Push(x); err != nil {
				return nil, err
			}
		}
	}

	return seen, nil
}

// recursive
func DFSr(g *ch09.Graph) []bool {
	n := g.Size()
	seen := make([]bool, n)
	for v := 0; v < n; v++ {
		dfsr(g, v, &seen)
	}
	return seen
}

func dfsr(g *ch09.Graph, v int, seen *[]bool) {
	(*seen)[v] = true

	for _, next := range g.VerticesConnectedWith(v) {
		if (*seen)[next] {
			continue
		}
		dfsr(g, next, seen)
	}
}

// by BFS
func Distance(g *ch09.Graph, s int) ([]int, error) {
	dist := make([]int, g.Size())
	for i := range dist {
		dist[i] = -1 // 「未訪問」で初期化
	}
	dist[s] = 0

	todo := ch08.NewQueue[int]()
	if err := todo.Push(s); err != nil {
		return nil, err
	}

	for !todo.IsEmpty() {
		v, err := todo.Pop()
		if err != nil {
			return nil, err
		}
		for _, x := range g.VerticesConnectedWith(v) {
			if dist[x] != -1 { // 探索済み
				continue
			}
			dist[x] = dist[v] + 1
			if err := todo.Push(x); err != nil {
				return nil, err
			}
		}

	}
	return dist, nil
}

// code 13.4
// DFS recursive
func PathExists(g *ch09.Graph, s, t int) bool {
	seen := make([]bool, g.Size())
	dfsr(g, s, &seen)
	return seen[t]
}

// code 13.5
// DFS recursive
func IsBipartite(g *ch09.Graph) bool {
	n := g.Size()
	color := make([]int, n) // 0: white, 1: black, -1: 未探索
	for i := range color {
		color[i] = -1 // 未探索
	}

	for v := 0; v < n; v++ {
		if color[v] != -1 {
			continue
		}
		if !isBipartite(g, v, 0, &color) {
			return false
		}
	}
	return true
}

// dfs, recursive
// cur: color of current node
func isBipartite(g *ch09.Graph, v, cur int, color *[]int) bool {
	(*color)[v] = cur
	for _, next := range g.VerticesConnectedWith(v) {
		if (*color)[next] != -1 { // 隣接頂点が既に色確定していた場合
			// 同じ色が隣接した場合は二部グラフではない
			if (*color)[next] == cur {
				return false
			}
		} else { // 隣接頂点の色が未確定の場合
			// 隣接頂点の色を変えて再起的に探索
			if !isBipartite(g, next, 1-cur, color) {
				return false
			}
		}
	}
	return true
}

// code 13.6
func TopologicalSort(g *ch09.Graph) []int {
	n := g.Size()
	seen := make([]bool, n)
	out := make([]int, 0)
	for v := 0; v < n; v++ {
		if seen[v] {
			continue
		}
		topologicalSort(g, v, &seen, &out)
	}

	reversed := make([]int, n)
	for i := range out {
		reversed[i] = out[n-1-i]
	}
	return reversed
}

// dfs
func topologicalSort(g *ch09.Graph, v int, seen *[]bool, out *[]int) {
	// fmt.Printf("%d-in\n", v)
	(*seen)[v] = true
	for _, next := range g.VerticesConnectedWith(v) {
		if (*seen)[next] {
			continue
		}
		topologicalSort(g, next, seen, out)
	}
	(*out) = append((*out), v)
	// fmt.Printf("%d-out\n", v)
}

// code 13.8
// treeはサイクルを持たない
func TreeSearchDepth(tree *ch09.Graph) []int {
	depth := make([]int, tree.Size())
	treeSearchDepth(tree, 0, 0, -1, &depth)
	return depth
}

// d: depth of node v. d = 0 if v is the root
// p: parent node of v. p = -1 if v is the root
func treeSearchDepth(tree *ch09.Graph, v, d, p int, depth *[]int) {
	(*depth)[v] = d
	for _, c := range tree.VerticesConnectedWith(v) {
		if c == p {
			continue // 親方向への探索を防ぐ
		}
		treeSearchDepth(tree, c, d+1, v, depth)
	}
}

// code 13.9
// treeはサイクルを持たない
func TreeSearchSubtreeSize(tree *ch09.Graph) []int {
	size := make([]int, tree.Size())
	treeSearchSubtreeSize(tree, 0, -1, &size)
	return size
}

func treeSearchSubtreeSize(tree *ch09.Graph, v, p int, size *[]int) {
	(*size)[v] = 1
	for _, c := range tree.VerticesConnectedWith(v) {
		if c == p {
			continue // 親方向への探索を防ぐ
		}
		treeSearchSubtreeSize(tree, c, v, size)
		(*size)[v] += (*size)[c]
	}
}

// ex 13.1
// DFS recursive
func CountConnectedGraph(g *ch09.Graph) int {
	var count int
	n := g.Size()
	seen := make([]bool, n)
	for v := 0; v < n; v++ {
		if seen[v] {
			continue
		}
		countConnectedGraph(g, v, &seen)
		count++
	}
	return count
}

func countConnectedGraph(g *ch09.Graph, v int, seen *[]bool) {
	(*seen)[v] = true
	for _, next := range g.VerticesConnectedWith(v) {
		if (*seen)[next] {
			continue
		}
		countConnectedGraph(g, next, seen)
	}
}

// ex 13.2
func PathExistsBFS(g *ch09.Graph, s, t int) (bool, error) {
	// dist, err := Distance(g, s)
	// if err != nil {
	// 	return false, err
	// }
	// return dist[t] != -1, nil
	// 上記でも結果同じ

	seen := make([]bool, g.Size())
	todo := ch08.NewQueue[int]()
	if err := todo.Push(s); err != nil {
		return false, err
	}

	for !todo.IsEmpty() {
		v, err := todo.Pop()
		if err != nil {
			return false, err
		}
		for _, next := range g.VerticesConnectedWith(v) {
			if seen[next] {
				continue
			}
			seen[next] = true
			if err := todo.Push(next); err != nil {
				return false, err
			}
		}

	}
	return seen[t], nil
}

// ex 13.3
func IsBipartiteBFS(g *ch09.Graph) (bool, error) {
	n := g.Size()
	color := make([]int, n) // black: 1, white: 0, 未探索: -1
	for i := range color {
		color[i] = -1
	}

	todo := ch08.NewQueue[int]()
	for v := 0; v < n; v++ {
		if color[v] != -1 {
			continue
		}
		result, err := isBipartiteBFS(g, todo, v, &color)
		if err != nil {
			return false, err
		}
		if !result {
			return false, nil
		}
	}
	return true, nil
}

func isBipartiteBFS(g *ch09.Graph, todo *ch08.Queue[int], s int, color *[]int) (bool, error) {
	(*color)[s] = 0
	if err := todo.Push(s); err != nil {
		return false, err
	}

	for !todo.IsEmpty() {
		v, err := todo.Pop()
		if err != nil {
			return false, err
		}
		for _, next := range g.VerticesConnectedWith(v) {
			if (*color)[next] != -1 { // 隣接頂点が色確定している場合
				if (*color)[next] == (*color)[v] {
					return false, nil
				}
				continue
			}
			// 隣接頂点の色が未確定の場合
			(*color)[next] = 1 - (*color)[v] // 逆の色にする
			if err := todo.Push(next); err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

// ex 13.4
// BFS
func ShortestPathOfMaze(h, w int, maze string) (int, error) {
	// field: h*w
	field := make([][]rune, h)
	for i := range field {
		field[i] = make([]rune, w)
	}

	sx, sy, gx, gy := -1, -1, -1, -1 // coordinate of start & goal
	maze = strings.Trim(maze, "\n")
	for i, line := range strings.Split(maze, "\n") {
		for j, c := range line {
			field[i][j] = c
			if c == 'S' {
				sx, sy = i, j
			} else if c == 'G' {
				gx, gy = i, j
			}
		}
	}

	type node [2]int
	que := ch08.NewQueue[node]()
	if err := que.Push(node{sx, sy}); err != nil {
		return 0, err
	}

	dist := make([][]int, h) // dist[i][j] は(i,j)への最短路長
	for i := range dist {
		dist[i] = make([]int, w)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[sx][sy] = 0

	// 上下左右の動きを表すベクトル
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for !que.IsEmpty() {
		cur, err := que.Pop()
		if err != nil {
			return 0, err
		}
		x, y := cur[0], cur[1]
		// 隣接頂点（上下左右）を順に見ていく
		for _, dir := range directions {
			nx := x + dir[0]
			ny := y + dir[1]

			// field外はスキップ
			if nx < 0 || nx >= h || ny < 0 || ny >= w {
				continue
			}
			// 壁内はスキップ
			if field[nx][ny] == '#' {
				continue
			}

			// 次のnodeが未探索ならば
			if dist[nx][ny] == -1 {
				dist[nx][ny] = dist[x][y] + 1 // 次nodeへの距離 = 現在のnodeへの距離 + 1
				if err := que.Push(node{nx, ny}); err != nil {
					return 0, err
				}
			}
		}
	}

	// for i := range dist {
	// 	fmt.Println(dist[i])
	// }

	return dist[gx][gy], nil
}
