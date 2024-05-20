package main

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, val := range nums {
		m[val] = struct{}{}
	}

	r := 0
	for n := range m {
		if _, exists := m[n-1]; !exists {
			i := 1
			for {
				if _, exists := m[n+i]; !exists {
					break
				}
				i++
			}
			if i > r {
				r = i
			}
		}
	}

	return r
}
