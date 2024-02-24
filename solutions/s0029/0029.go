package main

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Divide(dividend int, divisor int) int {
	isPositive := (dividend < 0) == (divisor < 0)

	dividend = abs(dividend)
	divisor = abs(divisor)
	fraction := 0

	for dividend >= divisor {
		bits := 0
		for dividend >= (divisor << (bits + 1)) {
			bits++
		}

		fraction += 1 << bits
		dividend -= divisor << bits
	}

	if !isPositive {
		fraction = -fraction
	}

	if fraction > math.MaxInt32 {
		return math.MaxInt32
	}

	return fraction
}
