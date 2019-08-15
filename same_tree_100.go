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
	root1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	root2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	fmt.Println(isSameTree(root1, root2))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q == nil {
		return false
	} else if p == nil && q != nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

