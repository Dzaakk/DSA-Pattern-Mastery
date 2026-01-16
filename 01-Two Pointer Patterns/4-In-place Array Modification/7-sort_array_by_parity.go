package main

func sortArrayByParity(nums []int) []int {
	l, r := 0, len(nums)-1

	for l < r {
		if nums[l]%2 == 0 {
			l++
			continue
		}
		if nums[r]%2 == 1 {
			r--
			continue
		}

		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

	return nums
}
