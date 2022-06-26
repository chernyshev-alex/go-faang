package graphs

func max(a, b int) int {
	if a > b {
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
