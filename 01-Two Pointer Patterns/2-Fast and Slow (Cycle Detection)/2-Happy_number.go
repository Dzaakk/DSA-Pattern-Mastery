package main

func isHappy(n int) bool {
	slow := n
	fast := n

	for {
		slow = sumOfSquares(slow)
		fast = sumOfSquares(sumOfSquares(fast))

		if fast == 1 {
			return true
		}

		if slow == fast {
			return false
		}
	}
}

func sumOfSquares(n int) int {
	sum := 0

	for n > 0 {
		d := n % 10
		sum += d * d
		n /= 10
	}

	return sum
}
