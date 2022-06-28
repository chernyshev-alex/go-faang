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

// == https://leetcode.com/problems/knight-probability-in-chessboard/  ==

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
