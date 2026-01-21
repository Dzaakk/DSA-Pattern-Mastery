package main

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0
	res := []int{}

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			if len(res) == 0 || res[len(res)-1] != nums1[i] {
				res = append(res, nums1[i])
			}

			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}
