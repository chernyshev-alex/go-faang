package graphs

import (
	"math"

	"github.com/go-mci-faans/arr"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// --- Adjacency Lists ---

func traversal_bfs(m [][]int) []int {
	seen, values := make(map[int]bool), make([]int, 0)

	q := make([]int, 0)
	q = append(q, 0)

	for len(q) > 0 {
		var v int
		v, q = q[0], q[1:]
		values = append(values, v)
		seen[v] = true

		for n := range m[v] {
			if !seen[n] {
				q = append(q, n)
			}
		}
	}
	return values
}

func traversal_dfs(v int, m [][]int, values *[]int, seen *map[int]bool) {
	*values, (*seen)[v] = append(*values, v), true
	for _, n := range m[v] {
		if !(*seen)[n] {
			traversal_dfs(n, m, values, seen)
		}
	}
}

// --- Adjacency Matrix ---

func traversalBfsMatrix(m [][]int) []int {
	seen, values, q := make(map[int]bool), make([]int, 0), make([]int, 0)
	q = append(q, 0)
	for row := 0; len(q) > 0; {
		row, q = q[0], q[1:]
		values, seen[row] = append(values, row), true
		for col, connected := range m[row] {
			if _, ok := seen[col]; !ok && connected > 0 {
				q = append(q, col)
			}
		}
	}
	return values
}

func traversalDfsMatrix(v int, m [][]int, values *[]int, seen *map[int]bool) {
	*values, (*seen)[v] = append(*values, v), true
	for n, connected := range m[v] {
		if _, ok := (*seen)[n]; !ok && connected > 0 {
			traversalDfsMatrix(n, m, values, seen)
		}
	}
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {

	type tdfs func(int, [][]int) int
	var dfs tdfs

	dfs = func(nodeId int, emplsAdjList [][]int) int {
		if len(emplsAdjList[nodeId]) == 0 {
			return 0
		}
		max_time := 0
		for _, n := range emplsAdjList[nodeId] {
			max_time = max(max_time, dfs(n, emplsAdjList))
		}
		return max_time + informTime[nodeId]
	}

	// make adjacence list
	employees_adj_ls := make([][]int, len(manager))
	for r := range employees_adj_ls {
		employees_adj_ls[r] = make([]int, 0)
	}
	for emp_id, manager_id := range manager {
		if manager_id != -1 {
			employees_adj_ls[manager_id] = append(employees_adj_ls[manager_id], emp_id)
		}
	}
	return dfs(headID, employees_adj_ls)
}

// Topological sort =========

// Course Schedule
func canFinish(numCourses int, prerequisites [][]int) bool {
	in_degree := make([]int, numCourses)
	for _, row := range prerequisites {
		in_degree[row[0]] += 1
	}

	stack := make([]int, 0)
	for idx, v := range in_degree {
		if v == 0 {
			stack = append([]int{idx}, stack...) // push front
		}
	}

	count := 0
	for len(stack) > 0 {
		var v int
		v, stack = stack[0], stack[1:]
		count += 1
		for _, p := range prerequisites {
			if v == p[1] {
				in_degree[p[0]] -= 1
				if in_degree[p[0]] == 0 {
					stack = append([]int{p[0]}, stack...)
				}
			}
		}

	}
	return count == numCourses
}

// Requires Positive Weights, no cycles
func networkDelayTime_Dejkstra(times [][]int, n int, k int) int {
	type Pair struct{ to, w int }
	heapComparator := func(a, b interface{}) bool { return a.(int) < b.(int) }
	minHeap := arr.NewBinaryHeap(heapComparator)

	dist, adjLs := make([]int, n), make([][]Pair, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxUint16
		adjLs[i] = make([]Pair, 0)
	}

	from_idx := k - 1
	dist[from_idx] = 0
	minHeap.Push(from_idx)

	for _, t := range times {
		from, to, w := t[0], t[1], t[2]
		adjLs[from-1] = append(adjLs[from-1], Pair{to - 1, w})
	}

	for minHeap.Size() > 0 {
		v, _ := minHeap.Pop()
		v_idx := v.(int)
		for _, p := range adjLs[v_idx] {
			if dist[v_idx]+p.w < dist[p.to] {
				dist[p.to] = dist[v_idx] + p.w
				minHeap.Push(p.to)
			}
		}
	}

	max_element := 0
	for _, v := range dist {
		max_element = max(max_element, v)
	}
	if max_element == math.MaxUint16 {
		return -1
	}
	return max_element
}

// Bellman Ford for Negative Weights
func networkDelayTime_BellmanFord(times [][]int, n int, k int) int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxUint16
	}
	dist[k-1] = 0
	for i := 0; i < n-1; i++ {
		count := 0
		for _, t := range times {
			from, to, w := t[0], t[1], t[2]
			if dist[from-1]+w < dist[to-1] {
				dist[to-1] = dist[from-1] + w
				count += 1
			}
		}
		if count == 0 {
			break
		}
	}
	max_element := 0
	for _, v := range dist {
		max_element = max(max_element, v)
	}
	if max_element == math.MaxUint16 {
		return -1
	}
	return max_element
}
