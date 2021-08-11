package main

import "log"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)

	log.Println(matrix)
}

func rotate(matrix [][]int) {
	n := len(matrix)

	oldX, oldY := 0, 0
	newX, newY := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if oldX == i && oldY == j {
				matrix[newX][newY] = matrix[j][n-i-1]
			} else {
				matrix[i][j] = matrix[j][n-i-1]
			}
			oldX, oldY = j, n-i-1
			newX, newY = i, j

			log.Println(oldX, oldY, newX, newY)

		}
	}
}
