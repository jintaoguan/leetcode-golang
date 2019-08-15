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
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	root2 := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(leafSimilar(root1, root2))
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	seq1 := make([]int, 0)
	leafSimilarHelper(root1, &seq1)
	seq2 := make([]int, 0)
	leafSimilarHelper(root2, &seq2)
	if len(seq1) != len(seq2) {
		return false
	}
	for i := 0; i < len(seq1); i++ {
		if seq1[i] != seq2[i] {
			return false
		}
	}
	return true
}

func leafSimilarHelper(root *TreeNode, sequence *[]int) {
	if root.Left == nil && root.Right == nil {
		*sequence = append(*sequence, root.Val)
		return
	}
	if root.Left != nil {
		leafSimilarHelper(root.Left, sequence)
	}
	if root.Right != nil {
		leafSimilarHelper(root.Right, sequence)
	}
}

