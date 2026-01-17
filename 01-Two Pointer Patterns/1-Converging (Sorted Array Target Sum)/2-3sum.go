package main

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)

	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 {
			break
		}

		left := i + 1
		right := n - 1

		for left < right {
			s := nums[i] + nums[left] + nums[right]
			if s < 0 {
				left++
			} else if s > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return res
}
