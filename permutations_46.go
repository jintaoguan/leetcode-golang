package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	permutations := make([][]int, 0)
	permuteHelper(nums, 0, &permutations)
	return permutations
}

func permuteHelper(nums []int, index int, result *[][]int) {
	if index >= len(nums) {
		permutation := make([]int, len(nums))
		copy(permutation, nums)
		*result = append(*result, permutation)
		return
	}
	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		permuteHelper(nums, index+1, result)
		nums[i], nums[index] = nums[index], nums[i]
	}
}

