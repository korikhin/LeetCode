package main

func findMin(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		m := i + (j-i)/2
		if nums[j] < nums[m] {
			i = m + 1
		} else {
			j = m
		}
	}

	return nums[j]
}
