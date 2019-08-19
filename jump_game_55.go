package main

import "fmt"

func main() {
	input1 := []int{2, 3, 1, 1, 4}
	input2 := []int{3, 2, 1, 0, 4}
	input3 := []int{0}
	fmt.Println(canJump(input1))
	fmt.Println(canJump(input2))
	fmt.Println(canJump(input3))
}

func canJump(nums []int) bool {
	maxPos := nums[0]
	for i := 0; i <= maxPos; i++ {
		maxPos = max(maxPos, i+nums[i])
		if maxPos >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

