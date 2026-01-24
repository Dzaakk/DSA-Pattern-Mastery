package main

import "sort"

type pair struct {
	val int
	cnt int
}

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	result := make([]int, 0, n-k+1)

	for i := 0; i <= n-k; i++ {

		freq := make(map[int]int)

		for j := i; j < i+k; j++ {
			freq[nums[j]]++
		}

		arr := make([]pair, 0, len(freq))
		for v, c := range freq {
			arr = append(arr, pair{v, c})
		}

		sort.Slice(arr, func(i, j int) bool {
			if arr[i].cnt != arr[j].cnt {
				return arr[i].cnt > arr[j].cnt
			}
			return arr[i].val > arr[j].val
		})

		sum := 0

		for j := 0; j < x && j < len(arr); j++ {
			sum += arr[j].val * arr[j].cnt
		}
		result = append(result, sum)
	}

	return result
}
