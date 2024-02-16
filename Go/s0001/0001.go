package s0001

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{j, i}
		}
		m[n] = i
	}

	return []int{}
}
