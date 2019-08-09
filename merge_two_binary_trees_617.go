package main

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
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	mergeTrees(root1, root2)
	mergeTrees2(root1, root2)
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 != nil && t2 != nil {
		return &TreeNode{
			Val:   t1.Val + t2.Val,
			Left:  mergeTrees(t1.Left, t2.Left),
			Right: mergeTrees(t1.Right, t2.Right),
		}
	} else if t1 != nil {
		return &TreeNode{
			Val:   t1.Val,
			Left:  mergeTrees(t1.Left, nil),
			Right: mergeTrees(t1.Right, nil),
		}
	} else if t2 != nil {
		return &TreeNode{
			Val:   t2.Val,
			Left:  mergeTrees(nil, t2.Left),
			Right: mergeTrees(nil, t2.Right),
		}
	} else {
		return nil
	}
}

func mergeTrees2(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	} else if t1 == nil {
		return t2
	} else if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees2(t1.Left, t2.Left)
	t1.Right = mergeTrees2(t1.Right, t2.Right)
	return t1
}

