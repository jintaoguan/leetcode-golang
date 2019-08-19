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
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	fmt.Println(buildTree(preorder, inorder).Val)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	positions := make(map[int]int)
	for i, n := range inorder {
		positions[n] = i
	}
	return buildTreeHelper(preorder, 0, 0, len(preorder), positions)
}

func buildTreeHelper(preorder []int, index int, start int, length int, positions map[int]int) *TreeNode {
	if length == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[index]}
	leftLength := positions[preorder[index]] - start
	root.Left = buildTreeHelper(preorder, index+1, start, leftLength, positions)
	root.Right = buildTreeHelper(preorder, index+1+leftLength, start+leftLength+1, length-leftLength-1, positions)
	return root
}

