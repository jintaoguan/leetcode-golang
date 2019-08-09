package main

import "fmt"

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
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(diameterOfBinaryTree(root))
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, diameter := diameterOfBinaryTreeHelper(root)
	return diameter - 1
}

func diameterOfBinaryTreeHelper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftDepth, leftDiameter := diameterOfBinaryTreeHelper(root.Left)
	rightDepth, rightDiameter := diameterOfBinaryTreeHelper(root.Right)
	depth := max(leftDepth, rightDepth) + 1
	maxDiameter := max(leftDepth+rightDepth+1, max(leftDiameter, rightDiameter))
	return depth, maxDiameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

