package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	n := len(nums)

	if n < 3 {
		return 0
	}

	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}

			if sum == target {
				return sum
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closest

}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
