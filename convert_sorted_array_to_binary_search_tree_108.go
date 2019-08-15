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
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sortedArrayToBST(nums).Val)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return sortedArrayToBSTHelper(nums, 0, len(nums)-1)
}

func sortedArrayToBSTHelper(nums []int, low int, high int) *TreeNode {
	if low > high {
		return nil
	}
	mid := (high + low) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBSTHelper(nums, low, mid-1),
		Right: sortedArrayToBSTHelper(nums, mid+1, high),
	}
	return root
}

