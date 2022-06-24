package graphs

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
	seen, values := make(map[int]bool), make([]int, 0)
	q := make([]int, 0)
	q = append(q, 0)

	for len(q) > 0 {
		var v int
		v, q = q[0], q[1:]
		values = append(values, v)
		seen[v] = true

		for connected, v := range m[v] {
			if connected > 0 && !seen[v] {
				q = append(q, v)
			}
		}
	}
	return values
}

func traversalDfsMatrix(v int, m [][]int, values *[]int, seen *map[int]bool) {
	*values, (*seen)[v] = append(*values, v), true
	for n, connected := range m[v] {
		if connected > 0 {
			if _, ok := (*seen)[n]; !ok {
				traversalDfsMatrix(n, m, values, seen)
			}
		}
	}
}
