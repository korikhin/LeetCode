package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	d := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[d] = nums[i]
			d++
		}
	}

	return d
}
