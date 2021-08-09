package main

import (
	"log"
)

func main() {
	log.Println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))
}

func isValidSudoku(board [][]byte) bool {
	a := make([]map[byte]bool, 9, 9) // 横9
	b := make([]map[byte]bool, 9, 9) // 竖9
	c := make([]map[byte]bool, 9, 9) // 3x3
	for i := 0; i < 9; i++ {
		a[i] = make(map[byte]bool)
		b[i] = make(map[byte]bool)
		c[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := a[i][board[i][j]]; ok {
				return false
			} else {
				a[i][board[i][j]] = true
			}

			if _, ok := b[j][board[i][j]]; ok {
				return false
			} else {
				b[j][board[i][j]] = true
			}

			if _, ok := c[i/3*3+j/3][board[i][j]]; ok {
				return false
			} else {
				c[i/3*3+j/3][board[i][j]] = true
			}
		}
	}
	return true
}
