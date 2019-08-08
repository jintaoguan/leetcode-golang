package main

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
	bstToGst(root)
	// fmt.Println(bstToGstHelper(root, 0))
}

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	bstToGstHelper(root, 0)
	return root
}

func bstToGstHelper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	oldVal := root.Val

	rightSum := bstToGstHelper(root.Right, sum)
	root.Val = rightSum + sum + root.Val

	leftSum := bstToGstHelper(root.Left, rightSum+oldVal+sum)
	return leftSum + rightSum + oldVal
}

