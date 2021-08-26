package main

import "log"

func main() {
	log.Println(uniquePaths(20, 20))

}

func uniquePaths(m int, n int) int {
	array := make([][]int, m, m)
	for i := range array {
		array[i] = make([]int, n, n)
	}
	for i := 0; i < m; i++ {
		array[i][0] = 1
	}

	for j := 0; j < n; j++ {
		array[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			array[i][j] = sum(array, i, j)
		}
	}

	return array[m-1][n-1]
}

func sum(array [][]int, row, col int) int {
	return array[row][col-1] + array[row-1][col]
}
