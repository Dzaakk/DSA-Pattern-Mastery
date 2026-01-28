package main

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	i, j := 0, n-1
	idx := n - 1

	for i <= j {
		leftVal := nums[i] * nums[i]
		rightVal := nums[j] * nums[j]

		if leftVal > rightVal {
			res[idx] = leftVal
			i++
		} else {
			res[idx] = rightVal
			j--
		}

		idx--
	}

	return res
}
