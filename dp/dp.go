package dp

import (
	"golang.org/x/exp/constraints"
)

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

//
// (O)= NxN
//
func minCostClimbingStairsTopBottom(cost []int) int {
	type TMinCostFn func(n int) int
	var minCost TMinCostFn

	minCost = func(n int) int {
		if n < 0 {
			return 0
		} else if n == 0 || n == 1 {
			return cost[n]
		}
		return cost[n] + min(minCost(n-1), minCost(n-2))
	}
	ln := len(cost)
	return min(minCost(ln-1), minCost(ln-2))
}

// Memoization
func minCostClimbingStairsMemo(cost []int) int {
	ln := len(cost)
	switch ln {
	case 0:
		return 0
	case 1:
		return cost[0]
	}

	dp := make([]int, ln)
	for i := range dp {
		if i < 2 {
			dp[i] = cost[i]
		} else {
			dp[i] = cost[i] + min(dp[i-1], dp[i-2])
		}
	}
	return min(dp[ln-1], dp[ln-2])
}

// Memoization Opt
func minCostClimbingStairsMemoOpt(cost []int) int {
	ln := len(cost)
	switch ln {
	case 0:
		return 0
	case 1:
		return cost[0]
	}

	dp1, dp2 := cost[0], cost[1]
	for i := 2; i < ln; i++ {
		cc := cost[i] + min(dp1, dp2)
		dp1 = dp2
		dp2 = cc
	}
	return min(dp1, dp2)
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]uint16, len(text1)+1)
	for i := range dp {
		dp[i] = make([]uint16, len(text2)+1)
	}

	for i, b1 := range []byte(text1) {
		for j, b2 := range []byte(text2) {
			if b1 == b2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return int(dp[len(text1)][len(text2)])
}
