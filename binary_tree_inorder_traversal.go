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
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	fmt.Println(inorderTraversal(root))
}

func inorderTraversal(root *TreeNode) []int {
	elements := make([]int, 0)
	inorderTraversalHelper(root, &elements)
	return elements
}

func inorderTraversalHelper(root *TreeNode, elements *[]int) {
	if root == nil {
		return
	}
	inorderTraversalHelper(root.Left, elements)
	*elements = append(*elements, root.Val)
	inorderTraversalHelper(root.Right, elements)
}

