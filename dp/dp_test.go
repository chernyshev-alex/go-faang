package dp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/min-cost-climbing-stairs/
//
// You are given an integer array cost where cost[i] is the cost of ith step on a staircase.
// Once you pay the cost, you can either climb one or two steps.
// You can either start from the step with index 0, or the step with index 1.
// Return the minimum cost to reach the top of the floor.
//
func Test_MinCostClimbingStairsTopBottom(t *testing.T) {
	var ts = []struct {
		cost []int
		exp  int
	}{{[]int{10, 15, 20}, 15}, {[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := minCostClimbingStairsTopBottom(ts[i].cost)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

// Memoization
func Test_MinCostClimbingStairsTopBottomMemo(t *testing.T) {
	var ts = []struct {
		cost []int
		exp  int
	}{{[]int{10, 15, 20}, 15}, {[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := minCostClimbingStairsMemo(ts[i].cost)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

// Memoization Opt.
func Test_MinCostClimbingStairsTopBottomMemoOpt(t *testing.T) {
	var ts = []struct {
		cost []int
		exp  int
	}{{[]int{10, 15, 20}, 15}, {[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := minCostClimbingStairsMemoOpt(ts[i].cost)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

// == https://leetcode.com/problems/knight-probability-in-chessboard/
// On an n x n chessboard, a knight starts at the cell (row, column) and attempts to make exactly k moves.
// The rows and columns are 0-indexed, so the top-left cell is (0, 0), and the bottom-right cell is (n - 1, n - 1).
// A chess knight has eight possible moves it can make, as illustrated below.
// Each move is two cells in a cardinal direction, then one cell in an orthogonal direction.
// Each time the knight is to move, it chooses one of eight possible moves uniformly at random (even if the piece would go off the chessboard) and moves there.
// The knight continues moving until it has made exactly k moves or has moved off the chessboard.
// Return the probability that the knight remains on the board after it has stopped moving.

func Test_KnightProbability(t *testing.T) {
	var ts = []struct {
		n, k, row, column int
		exp               float64
	}{{1, 0, 0, 0, 1.0}, {3, 2, 0, 0, 0.06250}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := knightProbability(ts[i].n, ts[i].k, ts[i].row, ts[i].column)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

func Test_KnightProbabilityMemo(t *testing.T) {
	var ts = []struct {
		n, k, row, column int
		exp               float64
	}{{1, 0, 0, 0, 1.0}, {3, 2, 0, 0, 0.06250}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := knightProbabilityMemo(ts[i].n, ts[i].k, ts[i].row, ts[i].column)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

func Test_KnightProbabilityBottomUp(t *testing.T) {
	var ts = []struct {
		n, k, row, column int
		exp               float64
	}{{1, 0, 0, 0, 1.0}, {3, 2, 0, 0, 0.06250}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := knightProbabilityBottomUp(ts[i].n, ts[i].k, ts[i].row, ts[i].column)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

func Test_KnightProbabilityBottomUpOpt(t *testing.T) {
	var ts = []struct {
		n, k, row, column int
		exp               float64
	}{{1, 0, 0, 0, 1.0}, {3, 2, 0, 0, 0.06250}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := knightProbabilityBottomUpOpt(ts[i].n, ts[i].k, ts[i].row, ts[i].column)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

// Longest Common Subsequence
//
// https://leetcode.com/problems/longest-common-subsequence
// Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.
// A subsequence of a string is a new string generated from the original string with some characters
// (can be none) deleted without changing the relative order of the remaining characters.
// For example, "ace" is a subsequence of "abcde".
// A common subsequence of two strings is a subsequence that is common to both strings
//

func Test_LongestCommonSubsequence(t *testing.T) {
	var ts = []struct {
		s1, s2 string
		exp    int
	}{{"abcde", "ace", 3}, {"abc", "abc", 3}, {"abc", "def", 0}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := longestCommonSubsequence(ts[i].s1, ts[i].s2)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}
