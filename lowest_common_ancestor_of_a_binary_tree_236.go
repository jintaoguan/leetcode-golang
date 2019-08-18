package main

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	p := &TreeNode{Val: 7}
	q := &TreeNode{Val: 4}
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  p,
				Right: q,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}
	fmt.Println(lowestCommonAncestor(root, p, q).Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if subtree(root.Left, p) && subtree(root.Left, q) {
		return lowestCommonAncestor(root.Left, p, q)
	} else if subtree(root.Right, p) && subtree(root.Right, q) {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

func subtree(root, target *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == target {
		return true
	}
	return subtree(root.Left, target) || subtree(root.Right, target)
}

