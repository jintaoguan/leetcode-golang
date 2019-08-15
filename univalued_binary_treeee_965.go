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
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(isUnivalTree(root))
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isUnivalTreeHelper(root, root.Val)
}

func isUnivalTreeHelper(root *TreeNode, value int) bool {
	if root == nil {
		return true
	}
	if root.Val != value {
		return false
	}
	return isUnivalTreeHelper(root.Left, root.Val) && isUnivalTreeHelper(root.Right, root.Val)
}

