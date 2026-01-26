package main

func reverseStr(s string, k int) string {
	b := []byte(s)

	for i := 0; i < len(b); i += 2 * k {
		l := i
		r := min(i+k-1, len(b)-1)

		for l < r {
			b[l], b[r] = b[r], b[l]
			l++
			r--
		}
	}

	return string(b)
}
