package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	result := make([]int, len(nums))
	result[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		result[i] = result[i+1] * nums[i+1]
	}
	product := 1
	for i := 0; i < len(nums); i++ {
		result[i] = result[i] * product
		product = product * nums[i]
	}
	return result
}

