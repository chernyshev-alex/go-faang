package arr

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TraversalDfs(t *testing.T) {
	const SZ = 3
	t.Run(t.Name(), func(t *testing.T) {
		m := makeMatrix(SZ, SZ, func(i, j int) int { return SZ*i + j })
		result := traversalDfs(m)
		assert.Equal(t, []int{0, 1, 2, 5, 8, 7, 4, 3, 6}, result)
	})
}

func Test_TraversalBfs(t *testing.T) {
	const SZ = 3
	t.Run(t.Name(), func(t *testing.T) {
		m := makeMatrix(SZ, SZ, func(i, j int) int { return SZ*i + j })
		result := traversalBfs(m)
		assert.Equal(t, []int{0, 1, 3, 2, 4, 6, 5, 7, 8}, result)
	})
}

func Test_DijkstraShortestPath(t *testing.T) {
	var ts = []struct {
		input            [][]Edge // adjacent list
		start, goal, exp uint
	}{{[][]Edge{
		{{node: 2, cost: 10}, {node: 1, cost: 1}},
		{{node: 3, cost: 2}},
		{{node: 1, cost: 2}, {node: 3, cost: 3}, {node: 4, cost: 1}},
		{{node: 0, cost: 7}, {node: 4, cost: 2}},
		{},
	}, 0, 4, 5}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			if result, ok := dijkstraShortestPath(ts[i].input, ts[i].start, ts[i].goal); ok {
				assert.Equal(t, result, ts[i].exp)
			} else {
				t.Errorf("not found path from %d to %d for input %v", ts[i].start, ts[i].goal, ts[i].input)
			}
		})
	}
}

// https://leetcode.com/problems/number-of-islands/
// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water),
// return the number of islands
//
func Test_NumIslands(t *testing.T) {
	var ts = []struct {
		input [][]byte
		exp   int
	}{{[][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}, 3}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := numIslands(ts[i].input)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

// https://leetcode.com/problems/rotting-oranges/
// You are given an m x n grid where each cell can have one of three values:
// 0 representing an empty cell,
// 1 representing a fresh orange, or
// 2 representing a rotten orange.
// Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
// Return the minimum number of minutes that must elapse until no cell has a fresh orange.
// If this is impossible, return -1.
////
func Test_OrangesRotting(t *testing.T) {
	var ts = []struct {
		input [][]int
		exp   int
	}{{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
		{[][]int{{0, 2}}, 0}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := orangesRotting(ts[i].input)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

// Given 2d array containing  -1 - walls, 0-gates, INF - empty room.
// Fill each empty room with a number of steps to the nearest gate.
// Leave INF if it is impossible to reach a gate.
//
func Test_WallsAndGates(t *testing.T) {
	const INF = math.MaxInt16
	var ts = []struct {
		input [][]int
		exp   [][]int
	}{{[][]int{{INF, -1, 0, INF}, {INF, INF, INF, 0}, {INF, -1, INF, -1}, {0, -1, INF, INF}},
		[][]int{{3, -1, 0, 1}, {2, 2, 1, 0}, {1, -1, 2, -1}, {0, -1, 3, 4}}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			wallsAndGates(ts[i].input)
			assert.Equal(t, ts[i].exp, ts[i].input)
		})
	}
}

// https://leetcode.com/problems/transpose-matrix/
// 	Given a 2D integer array matrix, return the transpose of matrix.
// The transpose of a matrix is the matrix flipped over its main diagonal,
// switching the matrix's row and column indices.
func Test_TransposeMatrix(t *testing.T) {
	var ts = []struct {
		input, exp [][]int
	}{{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{1, 4}, {2, 5}, {3, 6}}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := transpose(ts[i].input)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

// https://leetcode.com/problems/set-matrix-zeroes/
//
func Test_SetMatrixZeroes(t *testing.T) {
	var ts = []struct {
		input, exp [][]int
	}{{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			setZeroes(ts[i].input)
			assert.Equal(t, ts[i].exp, ts[i].input)
		})
	}
}

// https://leetcode.com/problems/search-a-2d-matrix/
//
func Test_Search_matrix(t *testing.T) {
	var ts = []struct {
		input  [][]int
		target int
		exp    bool
	}{{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 23, true}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := searchMatrix(ts[i].input, ts[i].target)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}
