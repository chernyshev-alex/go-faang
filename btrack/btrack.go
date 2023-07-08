package btrack

import (
	"bytes"
	"math"
)

func solveSudoku(board [][]byte) {

	ln := len(board)
	boxes, rows, cols := make([][]int, ln+1), make([][]int, ln+1), make([][]int, ln+1)
	for i := 0; i < ln; i++ {
		boxes[i], rows[i], cols[i] = make([]int, ln+1), make([]int, ln+1), make([]int, ln+1)
	}

	getBoxId := func(r, c int) int {
		ri, ci := math.Floor(float64(r)/3.0)*3.0, math.Floor(float64(c)/3.0)
		return int(ri + ci)
	}

	isValid := func(boxes, rows, cols int) bool {
		return boxes+rows+cols == 0
	}

	type FnSolveSudokuBackTracking func(r, c int) bool
	var solveSudokuBackTracking FnSolveSudokuBackTracking

	solveSudokuBackTracking = func(r, c int) bool {
		if r == len(board) || c == len(board[0]) {
			return true
		}
		if board[r][c] == '.' {
			for num := 1; num <= 9; num++ {
				board[r][c] = byte(num) + byte('0')
				boxId := getBoxId(r, c)
				if isValid(boxes[boxId][num], rows[r][num], cols[c][num]) {

					boxes[boxId][num], rows[r][num], cols[c][num] = 1, 1, 1

					if c == len(board[0])-1 {
						if solveSudokuBackTracking(r+1, 0) {
							return true
						}
					} else {
						if solveSudokuBackTracking(r, c+1) {
							return true
						}
					}

					boxes[boxId][num], rows[r][num], cols[c][num] = 0, 0, 0
				}
				board[r][c] = '.'
			}
		} else {
			if c == len(board[0])-1 {
				if solveSudokuBackTracking(r+1, 0) {
					return true
				}
			} else {
				if solveSudokuBackTracking(r, c+1) {
					return true
				}
			}
		}
		return false
	}

	for r := 0; r < ln; r++ {
		for c := 0; c < ln; c++ {
			if board[r][c] != '.' {
				n := board[r][c] - '0'
				boxId := getBoxId(r, c)
				boxes[boxId][n], rows[r][n], cols[c][n] = 1, 1, 1
			}
		}
	}
	solveSudokuBackTracking(0, 0)
}

// ---- Palindrome Partitioning --

func partition(s string) [][]string {
	isPalindrome := func(from, to int, str string) bool {
		for ; from < to; from, to = from+1, to-1 {
			if str[from] != str[to] {
				return false
			}
		}
		return true
	}

	type PartitioningFn func(int, []string)
	var doPartition PartitioningFn

	result, ln := make([][]string, 0), len(s)
	doPartition = func(index int, ps []string) {
		if index == ln {
			cps := make([]string, len(ps))
			copy(cps, ps)
			result = append(result, cps)
		} else {
			for i := index; i < ln; i++ {
				if isPalindrome(index, i, s) {
					part := s[index : i+1]
					ps = append(ps, part)
					doPartition(i+1, ps)
					ps = ps[:len(ps)-1]
				}
			}
		}
	}
	ps := make([]string, 0)
	doPartition(0, ps)
	return result
}

func solveNQueens(n int) [][]string {

	result := make([][]string, 0)

	isValid := func(queen_in_cols []int) bool {
		if len(queen_in_cols) == 1 {
			return true
		}
		r := len(queen_in_cols) - 1
		c := queen_in_cols[r]
		for i := 0; i < r; i++ {
			if queen_in_cols[i] == c || r-i == int(math.Abs(float64(queen_in_cols[i])-float64(c))) {
				return false
			}
		}
		return true
	}

	genSolution := func(n int, queen_in_cols []int) []string {
		sol := make([]string, 0)
		for r := range queen_in_cols {
			var b bytes.Buffer
			for c := 0; c < n; c++ {
				ss := '.'
				if c == queen_in_cols[r] {
					ss = 'Q'
				}
				b.WriteRune(ss)
			}
			sol = append(sol, b.String())
		}
		return sol
	}

	type SolveFnType func(row int, queen_in_cols []int)
	var doSolveFn SolveFnType

	doSolveFn = func(row int, queen_in_cols []int) {
		if row == n {
			sol := genSolution(n, queen_in_cols)
			result = append(result, sol)
		} else {
			for col := 0; col < n; col++ {
				queen_in_cols = append(queen_in_cols, col)
				if isValid(queen_in_cols) {
					doSolveFn(row+1, queen_in_cols)
				}
				queen_in_cols = queen_in_cols[:len(queen_in_cols)-1]
			}
		}
	}

	doSolveFn(0, make([]int, 0))
	return result
}

// String all permutations
func do_permute(s []rune, l, r int, result *[]string) {
	if l == r {
		*result = append(*result, string(s))
	} else {
		for i := l; i <= r; i++ {
			s[l], s[i] = s[i], s[l]
			do_permute(s, l+1, r, result)
			s[l], s[i] = s[i], s[l]
		}
	}
}

func permute(s string) []string {
	result := make([]string, 0)
	do_permute([]rune(s), 0, len(s)-1, &result)
	return result

}
