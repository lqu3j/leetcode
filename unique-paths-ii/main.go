package main

import "log"

func main() {
	array := [][]int{
		{0, 0},
		{1, 1},
		{0, 0},
	}
	log.Println(uniquePathsWithObstacles(array))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	array := make([][]int, m, m)
	for i := 0; i < m; i++ {
		array[i] = make([]int, n, n)
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			array[i][0] = 0
		} else {
			if i == 0 {
				array[0][0] = 1
			} else {
				array[i][0] = array[i-1][0]
			}
		}
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			array[0][j] = 0
		} else {
			if j == 0 {
				array[0][j] = 1
			} else {
				array[0][j] = array[0][j-1]
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				array[i][j] = sum(array, i, j)
			}
		}
	}

	return array[m-1][n-1]
}

func sum(array [][]int, row, col int) int {
	return array[row][col-1] + array[row-1][col]
}
