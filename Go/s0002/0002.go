package s0002

// Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{Val: 0}
	head := dummy

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		digit := sum % 10

		head.Next = &ListNode{Val: digit}
		head = head.Next
	}

	return dummy.Next
}
