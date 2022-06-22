package arr

import (
	"fmt"
	"math"
)

func makeMatrix(rows, cols int, fncell func(r, c int) int) (m [][]int) {
	m = make([][]int, rows)
	for r := range m {
		m[r] = make([]int, cols)
		if fncell != nil {
			for c := range m[r] {
				m[r][c] = fncell(r, c)
			}
		}
	}
	return
}

func directions() [][]int {
	var DIRECTIONS = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	return DIRECTIONS
}

func dfs(m [][]int, r, c int, visited [][]int, path *[]int) {
	if r < 0 || c < 0 || r >= len(m) || c >= len(m[0]) || visited[r][c] == 1 {
		return
	}
	visited[r][c] = 1
	*path = append(*path, m[r][c])
	for _, d := range directions() {
		dfs(m, r+d[0], c+d[1], visited, path)
	}
}

func traversalDfs(m [][]int) (path []int) {
	visited := makeMatrix(len(m), len(m[0]), nil)
	path = make([]int, 0)
	dfs(m, 0, 0, visited, &path)
	return path
}

type Pair struct {
	r, c int
}

func traversalBfs(m [][]int) (path []int) {
	visited := makeMatrix(len(m), len(m[0]), nil)
	path = make([]int, 0)

	q := append([]Pair(nil), Pair{0, 0})
	for len(q) > 0 {
		var e Pair
		e, q = q[0], q[1:]
		r, c := e.r, e.c

		if r < 0 || c < 0 || r >= len(m) || c >= len(m[0]) || visited[r][c] == 1 {
			continue
		}

		visited[r][c] = 1
		path = append(path, m[r][c])
		for _, d := range directions() {
			q = append(q, Pair{r: r + d[0], c: c + d[1]})
		}
	}
	return
}

type Edge struct {
	node, cost uint
}

type State struct {
	cost, position uint
}

func dijkstraShortestPath(als [][]Edge, start, goal uint) (uint, bool) {
	dist := make([]uint, len(als))
	for i := range dist {
		dist[i] = math.MaxUint32
	}

	fCompareState := func(a, b interface{}) bool { return a.(State).cost < b.(State).cost }
	minHeap := NewBinaryHeap(fCompareState)
	minHeap.Push(State{0, start})
	dist[start] = 0

	for v, ok := minHeap.Pop(); ok; v, ok = minHeap.Pop() {
		var s State = v.(State)

		if s.position == goal {
			return s.cost, true
		}
		if s.cost > dist[s.position] {
			continue
		}

		for _, edge := range als[s.position] {
			next := State{
				cost:     s.cost + edge.cost,
				position: edge.node,
			}
			if next.cost < dist[next.position] {
				minHeap.Push(next)
				dist[next.position] = next.cost
			}
		}
	}
	return 0, false
}

func numIslands(grid [][]byte) int {
	result := 0
	if len(grid) == 0 {
		return result
	}
	q := make([]Pair, 0)
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '1' {
				result += 1

				grid[row][col] = '0'
				q = append(q, Pair{r: row, c: col})

				for len(q) > 0 {
					var v Pair
					v, q = q[0], q[1:]
					for _, d := range directions() {
						nextRow, nextCol := v.r+d[0], v.c+d[1]
						if nextRow < 0 || nextCol < 0 ||
							nextRow >= len(grid) || nextCol >= len(grid[0]) {
							continue
						}

						if grid[nextRow][nextCol] == '1' {
							q = append(q, Pair{r: nextRow, c: nextCol})
							grid[nextRow][nextCol] = '0'
						}
					}
				}
			}
		}
	}
	return result
}

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	freshOranges := 0
	rotten := make([]Pair, 0)
	for row := range grid {
		for col := range grid[row] {
			switch grid[row][col] {
			case 1:
				freshOranges += 1
			case 2:
				rotten = append(rotten, Pair{r: row, c: col})
			}
		}
	}

	mins := 0
	for currentQueueSize := len(rotten); len(rotten) > 0; {
		if currentQueueSize == 0 {
			currentQueueSize = len(rotten)
			mins += 1
		}

		var v Pair
		v, rotten = rotten[0], rotten[1:]

		currentQueueSize -= 1

		for _, d := range directions() {
			nextRow, nextCol := v.r+d[0], v.c+d[1]
			if nextRow < 0 || nextCol < 0 ||
				nextRow >= len(grid) || nextCol >= len(grid[0]) {
				continue
			}

			if grid[nextRow][nextCol] == 1 {
				grid[nextRow][nextCol] = 2
				freshOranges -= 1
				rotten = append(rotten, Pair{r: nextRow, c: nextCol})

			}
		}
	}
	if freshOranges != 0 {
		return -1
	}
	return mins
}

func wallsAndGates(grid [][]int) {
	var fnDfs func(r, c, cur_value int)

	fnDfs = func(r, c, cur_value int) {
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[r]) || cur_value > grid[r][c] {
			return
		}
		grid[r][c] = cur_value
		for _, d := range directions() {
			fnDfs(r+d[0], c+d[1], cur_value+1)
		}
	}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 0 {
				fnDfs(r, c, 0)
			}
		}
	}
}

func transpose(matrix [][]int) [][]int {
	tr := make([][]int, len(matrix[0]))
	for r := range tr {
		tr[r] = make([]int, len(matrix))
	}
	for r := range matrix {
		for c := range matrix[r] {
			fmt.Printf("r%d ; c%d\n", r, c)
			tr[c][r] = matrix[r][c]
		}
	}
	return tr
}

func setZeroes(matrix [][]int) [][]int {
	panic("not implemented")

}

func searchMatrix(matrix [][]int, target int) bool {
	panic("not implemented")
}
