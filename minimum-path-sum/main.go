package main

import "log"

func main() {
	array := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	min := minPathSum(array)
	log.Println(min)
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	array := make([][]int, m, m)
	for i := range array {
		array[i] = make([]int, n, n)
	}
	for i := 0; i < m; i++ {
		if i == 0 {
			array[i][0] = grid[i][0]
		} else {
			array[i][0] = array[i-1][0] + grid[i][0]
		}
	}
	for j := 0; j < n; j++ {
		if j == 0 {
			array[0][j] = grid[0][j]
		} else {
			array[0][j] = array[0][j-1] + grid[0][j]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			array[i][j] = grid[i][j] + min(array, i, j)
		}
	}
	return array[m-1][n-1]
}

func min(array [][]int, x, y int) int {
	if array[x][y-1] > array[x-1][y] {
		return array[x-1][y]
	} else {
		return array[x][y-1]
	}
}
