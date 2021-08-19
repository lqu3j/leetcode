package main

import (
	"log"
)

func main() {
	log.Printf("%q", solveNQueens(8))
}

func solveNQueens(n int) [][]string {
	solves := [][]string{}

	for k := 0; k < n; k++ {
		i, j := 0, k
		solve := make([]string, n, n)
		for i >= 0 && j < n {
			if isValid(solve, i, j) {
				solve[i] = initRow(n, j)
				if i == n-1 {
					solves = append(solves, solve)
					solve = docopy(solve, n)
					solve[i] = ""
					j++
					i, j = calc(solve, i, j)

				} else {
					i++
					j = 0
				}
			} else {
				j++
				i, j = calc(solve, i, j)
			}
		}
	}
	return solves
}

func docopy(solve []string, n int) []string {
	tmp := make([]string, n, n)
	copy(tmp, solve)
	return tmp
}

func calc(solve []string, i, j int) (int, int) {
	for j >= len(solve) {
		i--
		if i <= 0 {
			break
		}
		j = getLastIndex(solve, i) + 1
		solve[i] = ""
	}
	return i, j
}

func initRow(n, column int) string {
	row := make([]byte, n, n)
	for i := range row {
		if i == column {
			row[i] = 'Q'
		} else {
			row[i] = '.'
		}
	}
	return string(row)
}

func getLastIndex(solve []string, row int) int {
	for j := range solve[row] {
		if solve[row][j] == 'Q' {
			return j
		}
	}
	return 0
}

func isValid(solve []string, row, column int) bool {
	if row == 0 {
		return true
	}

	for i := 0; i < row; i++ {
		if solve[i][column] == 'Q' {
			return false
		}
	}
	// 左上角
	x, y := row-1, column-1
	for x >= 0 && y >= 0 {
		if solve[x][y] == 'Q' {
			return false
		}
		x--
		y--
	}
	// 右上角
	x, y = row-1, column+1
	for x >= 0 && y < len(solve) {
		if solve[x][y] == 'Q' {
			return false
		}
		x--
		y++
	}

	return true
}
