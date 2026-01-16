package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	top, bottom := 0, n-1
	left, right := 0, n-1

	num := 1
	max := n * n

	for num <= max {

		// left to right (top row)
		for col := left; col <= right && num <= max; col++ {
			matrix[top][col] = num
			num++
		}
		top++

		// top to bottom (right column)
		for row := top; row <= bottom && num <= max; row++ {
			matrix[row][right] = num
			num++
		}
		right--

		// right to left (bottom row)
		for col := right; col >= left && num <= max; col-- {
			matrix[bottom][col] = num
			num++
		}
		bottom--

		// bottom to top (left column)
		for row := bottom; row >= top && num <= max; row-- {
			matrix[row][left] = num
			num++
		}
		left++
	}

	return matrix
}
