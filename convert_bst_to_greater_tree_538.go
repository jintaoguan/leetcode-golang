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
				Val: 0,
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}
	fmt.Println(convertBST(root).Val)
}

func convertBST(root *TreeNode) *TreeNode {
	convertBSTHelper(root, 0)
	return root
}

func convertBSTHelper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	oldVal := root.Val
	root.Val += sum
	rightSum := convertBSTHelper(root.Right, sum)
	root.Val += rightSum
	leftSum := convertBSTHelper(root.Left, sum+rightSum+oldVal)
	return leftSum + rightSum + oldVal
}

