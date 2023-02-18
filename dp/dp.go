package dp

import (
	"strings"

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

func knight_directions() [][]int {
	return [][]int{{-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}}
}

func knightProbability(n int, k int, row int, column int) float64 {
	if row < 0 || column < 0 || row >= n || column >= n {
		return 0.0
	}
	if k == 0 {
		return 1.0
	}
	res := 0.0
	for _, d := range knight_directions() {
		res += knightProbability(n, k-1, row+d[0], column+d[1]) / 8.0
	}
	return res
}

func knightProbabilityMemo(n int, k int, row int, column int) float64 {

	dp := make([][][]float64, k+1)
	for i := range dp {
		dp[i] = make([][]float64, n)
		for j := range dp[i] {
			dp[i][j] = make([]float64, n)
			for l := range dp[i][j] {
				dp[i][j][l] = -1
			}
		}
	}

	type TSolveFun func(n int, k int, row int, column int) float64
	var solveFun TSolveFun
	solveFun = func(n int, k int, row int, col int) float64 {
		if row < 0 || col < 0 || row >= n || col >= n {
			return 0.0
		}
		if k == 0 {
			return 1.0
		}
		if dp[k][row][col] != -1.0 {
			return dp[k][row][col]
		}

		res := 0.0
		for _, d := range knight_directions() {
			res += solveFun(n, k-1, row+d[0], col+d[1]) / 8.0
		}
		dp[k][row][col] = res
		return res
	}
	return solveFun(n, k, row, column)
}

func knightProbabilityBottomUp(n int, k int, row int, column int) float64 {
	dp := make([][][]float64, k+1)
	for i := range dp {
		dp[i] = make([][]float64, n)
		for j := range dp[i] {
			dp[i][j] = make([]float64, n)
			for l := range dp[i][j] {
				dp[i][j][l] = 1.0
			}
		}
	}
	for step := 1; step <= k; step++ {
		for row := 0; row < n; row++ {
			for col := 0; col < n; col++ {
				for _, d := range knight_directions() {
					prev_row, prev_col := row+d[0], col+d[1]
					if prev_row >= 0 && prev_row < n && prev_col >= 0 && prev_col < n {
						dp[step][row][col] += dp[(step - 1)][prev_row][prev_col] / 8.0
					}
				}
			}
		}
	}
	res := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += dp[k][i][j]
		}
	}
	return res
}

func knightProbabilityBottomUpOpt(n int, k int, row int, column int) float64 {
	dp_prev, dp_next := make([][]float64, n), make([][]float64, n)
	for r := 0; r < n; r++ {
		dp_prev[r] = make([]float64, n)
		dp_next[r] = make([]float64, n)
	}

	dp_prev[row][column] = 1.0
	for step := 1; step <= k; step++ {
		for row := 0; row < n; row++ {
			for col := 0; col < n; col++ {
				for _, d := range knight_directions() {
					prev_row, prev_col := row+d[0], col+d[1]
					if prev_row >= 0 && prev_row < n && prev_col >= 0 && prev_col < n {
						dp_next[row][col] += dp_prev[prev_row][prev_col] / 8.0
					}
				}
			}
		}
		dp_prev = dp_next
		dp_next = make([][]float64, n)
		for r := range dp_next {
			dp_next[r] = make([]float64, n)
		}
	}

	res := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += dp_prev[i][j]
		}
	}
	return res
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

//"abcdxyz", "xyzabcd", "abcd"
func longestCommonSubStr(s1, s2 string) string {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}
	lcsLen, lcsEnd := 0, 0
	for i1 := 1; i1 <= len(s1); i1++ {
		for i2 := 1; i2 <= len(s2); i2++ {
			if s1[i1-1] == s2[i2-1] {
				dp[i1][i2] = dp[i1-1][i2-1] + 1
				if dp[i1][i2] > lcsLen {
					lcsLen = dp[i1][i2]
					lcsEnd = i1
				}
			}
		}
	}
	return s1[lcsEnd-lcsLen : lcsEnd]

}

func longestCommonSubInArray(ss []string) string {
	s := ss[0]
	// generate all substrings for ss[0]
	result := ""
	for i := range s {
		for j := i + 1; j <= len(s); j++ {
			//check if substring is common to all words
			m := 1
			for ; m < len(ss); m++ {
				if !strings.Contains(ss[m], s[i:j]) {
					break
				}
			}
			// all strings are contains substring
			if m == len(ss) && len(s[i:j]) > len(result) {
				result = s[i:j] // found  longer
			}
		}
	}
	return result
}

// {"AGGTAB", "GXTXAYB", "GTAB"}}
func longestCommonSubsequenceStr(s1 string, s2 string) string {
	result := ""
	for i := 0; i < len(s2); i++ {
		for j := 0; j < len(s1); j++ {
			if s2[i] == s1[j] {
				result += string(s2[i])
				break
			}
		}
	}
	return result
}

func getMaxGold(gold [][]int, n, m int) int {

	dp := make([][]int, n)
	for r := range dp {
		dp[r] = make([]int, m)
		for c := range dp[r] {
			dp[r][c] = -1
		}
	}

	type CollectGoldFn func(r, c int) int
	var collectGoldFn CollectGoldFn

	collectGoldFn = func(r, c int) int {
		if (r < 0) || (r == n) || (c == m) {
			return 0
		}

		rightUpperDiag := collectGoldFn(r-1, c+1)

		right := collectGoldFn(r, c+1)

		rightLowerDiag := collectGoldFn(r+1, c+1)

		return gold[r][c] + max(max(rightUpperDiag, rightLowerDiag), right)

	}

	maxGold := 0
	for row := 0; row < n; row++ {
		gCollected := collectGoldFn(row, 0)
		maxGold = max(maxGold, gCollected)
	}
	return maxGold

}
