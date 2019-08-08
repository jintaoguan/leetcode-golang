package main

import (
	"fmt"
)

// TreeNode strut
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val:  15,
			Left: nil,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}
	fmt.Println(rangeSumBST(root, 7, 15))
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}
	sum += rangeSumBST(root.Left, L, R)
	sum += rangeSumBST(root.Right, L, R)
	return sum
}

