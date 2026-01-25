package main

func characterReplacement(s string, k int) int {
	freq := make([]int, 26)
	left := 0
	maxFreq := 0
	result := 0

	for right := 0; right < len(s); right++ {
		idx := s[right] - 'A'
		freq[idx]++

		if freq[idx] > maxFreq {
			maxFreq = freq[idx]
		}

		if (right-left+1)-maxFreq > k {
			freq[s[left]-'A']--
			left++
		}

		if right-left+1 > result {
			result = right - left + 1
		}

	}

	return result
}
