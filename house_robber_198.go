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
	steal := make([]int, length)
	steal[0], steal[1] = nums[0], nums[1]
	nosteal := make([]int, length)
	nosteal[0], nosteal[1] = 0, nums[0]
	for i := 2; i < len(nums); i++ {
		steal[i] = max(steal[i-2]+nums[i], nosteal[i-1]+nums[i])
		nosteal[i] = max(steal[i-1], nosteal[i-1])
	}
	return max(steal[length-1], nosteal[length-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

