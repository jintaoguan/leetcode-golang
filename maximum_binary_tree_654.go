package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	array := []int{3, 2, 1, 6, 0, 5}
	constructMaximumBinaryTree(array)
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMaximumBinaryTreeHelper(nums, 0, len(nums)-1)
}

func constructMaximumBinaryTreeHelper(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	if left == right {
		return &TreeNode{Val: nums[left]}
	}
	maxElement := nums[left]
	rootIndex := left
	for i := left + 1; i <= right; i++ {
		if nums[i] > maxElement {
			rootIndex = i
			maxElement = nums[i]
		}
	}
	root := &TreeNode{Val: maxElement}
	root.Left = constructMaximumBinaryTreeHelper(nums, left, rootIndex-1)
	root.Right = constructMaximumBinaryTreeHelper(nums, rootIndex+1, right)
	return root
}

