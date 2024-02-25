package main

func LongestValidParentheses(s string) int {
	if s == "" {
		return 0
	}

	m := 0
	stack := []int{-1}

	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}

		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		cur := i - stack[len(stack)-1]
		if cur > m {
			m = cur
		}
	}

	return m
}
