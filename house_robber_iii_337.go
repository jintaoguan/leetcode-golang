package main

import (
	"fmt"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		},
	}
	fmt.Println(rob(root))
}

func rob(root *TreeNode) int {
	max1, max2 := robHelper(root)
	return max(max1, max2)
}

func robHelper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftMax1, leftMax2 := robHelper(root.Left)
	rightMax1, rightMax2 := robHelper(root.Right)
	steal := leftMax2 + rightMax2 + root.Val
	nosteal := max(leftMax1, leftMax2) + max(rightMax1, rightMax2)
	return steal, nosteal
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

