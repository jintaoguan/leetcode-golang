package main

import "fmt"

func main() {
	array := []int{2, 1, 1, 2}
	fmt.Println(rob(array))
}

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	} else if length == 1 {
		return nums[0]
	} else if length == 2 {
		return max(nums[0], nums[1])
	}
	values := make([]int, length)
	values[0] = nums[0]
	values[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		values[i] = max(values[i-1], values[i-2]+nums[i])
	}
	return values[length-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

