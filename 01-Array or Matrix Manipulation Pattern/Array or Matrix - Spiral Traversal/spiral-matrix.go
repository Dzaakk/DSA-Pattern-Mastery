package main

func spiralOrder(matrix [][]int) []int {
	result := []int{}

	m := len(matrix)
	n := len(matrix[0])

	top, bottom := 0, m-1
	left, right := 0, n-1

	for top <= bottom && left <= right {

		// left to right (top row)
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++

		// top to bottom (right column)
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right--

		// right to left (bottom row)
		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom--
		}

		// bottom to top (left column)
		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++
		}
	}

	return result
}
