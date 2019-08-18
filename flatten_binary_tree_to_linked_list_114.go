package main

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
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	flatten(root)
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	leftSubtree := root.Left
	rightSubtree := root.Right
	flatten(leftSubtree)
	flatten(rightSubtree)
	root.Left = nil
	root.Right = leftSubtree
	tail := root
	for tail.Right != nil {
		tail = tail.Right
	}
	tail.Right = rightSubtree
}

