package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricBranches(root.Left, root.Right)
}

// Same function from "100. Same Tree"
func isSymmetricBranches(a *TreeNode, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.Val != b.Val {
		return false
	}

	// ...but leafs are swapped
	return isSymmetricBranches(a.Left, b.Right) && isSymmetricBranches(a.Right, b.Left)
}
