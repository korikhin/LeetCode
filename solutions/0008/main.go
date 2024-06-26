package main

import "math"

func myAtoi(s string) int {
	i, n, a, sign := 0, len(s), 0, 1

	// Ignore whitespaces
	for i < n && s[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	}

	// Check the sign
	if s[i] == '-' || s[i] == '+' {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	for ; i < n && s[i] >= '0' && s[i] <= '9'; i++ {
		digit := int(s[i] - '0')

		// Check overflow before doing math
		if a > (math.MaxInt32-digit)/10 {
			if sign > 0 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		a = a*10 + digit
	}

	return a * sign
}
