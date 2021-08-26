package main

func main() {

}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n, n)
	}

	x, y := 0, 0
	left, right := 0, n-1
	top, bottom := 0, n-1
	value := 1
	max := n * n
	for top <= bottom && left <= right && value <= max {
		for y <= right {
			matrix[x][y] = value
			value++
			y++
		}
		top++
		y--
		x++

		for x <= bottom {
			matrix[x][y] = value
			value++
			x++
		}
		right--
		x--
		y--

		for y >= left {
			matrix[x][y] = value
			value++
			y--
		}
		bottom--
		y++
		x--

		for x >= top {
			matrix[x][y] = value
			x--
			value++
		}
		left++
		x++
		y++
	}
	return matrix
}
