package main

func containsNearbyDuplicate(nums []int, k int) bool {
	mapNum := make(map[int]int)

	for i, n := range nums {
		if idx, ok := mapNum[n]; ok {
			if abs(idx-i) <= k {
				return true
			}
		}

		mapNum[n] = i
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
