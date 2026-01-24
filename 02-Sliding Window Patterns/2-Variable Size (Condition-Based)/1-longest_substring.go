package main

func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[rune]int)
	left := 0
	maxLen := 0

	for right, ch := range []rune(s) {
		if prev, ok := lastSeen[ch]; ok && prev >= left {
			left = prev + 1
		}

		lastSeen[ch] = right

		currLen := right - left + 1
		if currLen > maxLen {
			maxLen = currLen
		}
	}

	return maxLen
}
