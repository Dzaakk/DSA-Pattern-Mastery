package main

func findDuplicate(nums []int) int {

	for i := range nums {

		idx := abs(nums[i]) - 1

		if nums[idx] > 0 {
			nums[idx] *= -1
		} else {
			return abs(nums[i])
		}
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
