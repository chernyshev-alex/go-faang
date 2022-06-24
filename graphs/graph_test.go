package graphs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// --- Adjacency Lists ---

func Test_TraversalBfs(t *testing.T) {
	var ts = []struct {
		input [][]int
		exp   []int
	}{{[][]int{{1, 3}, {0}, {3, 8}, {0, 2, 4, 5}, {3, 6}, {3}, {4, 7}, {6}, {2}}, []int{0, 1, 3, 2, 4, 5, 8, 6, 7}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := traversal_bfs(ts[i].input)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

func Test_TraversalDfs(t *testing.T) {
	var ts = []struct {
		input [][]int
		exp   []int
	}{{[][]int{{1, 3}, {0}, {3, 8}, {0, 2, 4, 5}, {3, 6}, {3}, {4, 7}, {6}, {2}}, []int{0, 1, 3, 2, 8, 4, 6, 7, 5}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			values := make([]int, 0)
			seen := make(map[int]bool)
			traversal_dfs(0, ts[i].input, &values, &seen)
			assert.Equal(t, ts[i].exp, values)
		})
	}
}

// --- Adjacency Matrix ---

func Test_TraversalBfsMatrix(t *testing.T) {
	var ts = []struct {
		input [][]int
		exp   []int
	}{{[][]int{
		{0, 1, 0, 1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0}}, []int{0, 1, 3, 2, 4, 5, 8, 6, 7}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := traversalBfsMatrix(ts[i].input)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

func Test_TraversalDfsMatrix(t *testing.T) {
	var ts = []struct {
		input [][]int
		exp   []int
	}{{[][]int{
		{0, 1, 0, 1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0}}, []int{0, 1, 3, 2, 8, 4, 6, 7, 5}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			values := make([]int, 0)
			seen := make(map[int]bool)
			traversalDfsMatrix(0, ts[i].input, &values, &seen)
			assert.Equal(t, ts[i].exp, values)
		})
	}
}
