package main

import "log"

func main() {
	// array := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	array := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	// array := [][]int{
	// 	{1, 2, 3, 10, 11},
	// 	{4, 5, 6, 12, 13},
	// 	{7, 8, 9, 14, 15},
	// }
	log.Println(spiralOrder(array))
}

func spiralOrder(matrix [][]int) []int {
	x, y := 0, 0
	left, right := 0, len(matrix[0])-1
	top, bottom := 0, len(matrix)-1
	length := len(matrix[0]) * len(matrix)
	result := make([]int, 0, len(matrix[0])*len(matrix))
	for top <= bottom && left <= right {
		for y <= right {
			result = append(result, matrix[x][y])
			y++
		}
		top++
		y--
		x++

		for x <= bottom {
			result = append(result, matrix[x][y])
			x++
		}
		right--
		x--
		y--

		for y >= left {
			result = append(result, matrix[x][y])
			y--
		}
		bottom--
		y++
		x--

		for x >= top {
			result = append(result, matrix[x][y])
			x--
		}
		left++
		x++
		y++
	}

	return result[:length]
}
