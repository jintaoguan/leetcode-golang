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
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 1,
			// Right: &TreeNode{
			// 	Val: 2,
			// },
		},
	}
	fmt.Println(isSymmetric(root))
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return isSymmetricHelper(t1.Right, t2.Left) && isSymmetricHelper(t1.Left, t2.Right)
}

