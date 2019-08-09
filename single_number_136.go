package main

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	array := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(array))
}

func singleNumber(nums []int) int {
	xor := 0
	for _, elem := range nums {
		xor ^= elem
	}
	return xor
}

