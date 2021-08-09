package main

func main() {

}

func solveSudoku(board [][]byte) {
	a, b, c := filed(board)

}

func filed(board [][]byte) ([]map[byte]bool, []map[byte]bool, []map[byte]bool) {
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
			a[i][board[i][j]] = true
			b[j][board[i][j]] = true
			c[i/3*3+j/3][board[i][j]] = true

		}
	}
	return a, b, c
}
