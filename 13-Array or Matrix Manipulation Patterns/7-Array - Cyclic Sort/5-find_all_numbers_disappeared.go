package main

func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	listNum := make([]int, len(nums))

	for _, n := range nums {
		x := n - 1
		listNum[x]++
	}

	for i, num := range listNum {
		if num == 0 {
			res = append(res, i+1)
		}
	}

	return res
}
