package btrack

import "math"

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
