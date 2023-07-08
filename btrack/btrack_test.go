package btrack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sudoku-solver
// Write a program to solve a Sudoku puzzle by filling the empty cells.
// A sudoku solution must satisfy all of the following rules:
// Each of the digits 1-9 must occur exactly once in each row.
// Each of the digits 1-9 must occur exactly once in each column.
// Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
// The '.' character indicates empty cells.
func Test_Sudoku(t *testing.T) {
	var ts = []struct {
		input [][]byte
		exp   [][]byte
	}{{
		[][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
		[][]byte{
			{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
			{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
			{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
			{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
			{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
			{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
			{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
			{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
			{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
		},
	},
	}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			solveSudoku(ts[i].input)
			assert.Equal(t, ts[i].exp, ts[i].input)
		})
	}
}

func Test_PalindromPartition(t *testing.T) {
	var ts = []struct {
		input string
		exp   [][]string
	}{{"cdd", [][]string{{"c", "d", "d"}, {"c", "dd"}}},
		{"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{"a", [][]string{{"a"}}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := partition(ts[i].input)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

// https://leetcode.com/problems/n-queens/
//
// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens
// attack each other.
// Given an integer n, return all distinct solutions to the n-queens puzzle.
// You may return the answer in any order.
// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.'
// both indicate a queen and an empty space, respectively.
func Test_SolveNQueens(t *testing.T) {
	var ts = []struct {
		input int
		exp   [][]string
	}{{1, [][]string{{"Q"}}},
		{4, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
	}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := solveNQueens(ts[i].input)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

// Print all permutations of a given string
func Test_Permutation(t *testing.T) {
	var ts = []struct {
		input string
		exp   []string
	}{{"ABC", []string{"ABC", "ACB", "BAC", "BCA", "CBA", "CAB"}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := permute(ts[i].input)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}
