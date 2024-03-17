package main

func HammingWeight(num uint32) int {
	var w uint8
	for num > 0 {
		num &= num - 1
		w++
	}

	return int(w)
}
