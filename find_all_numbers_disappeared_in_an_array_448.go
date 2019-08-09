package main

import "fmt"

func main() {
	array := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(array))
}

func findDisappearedNumbers(nums []int) []int {
	index := 0
	for index < len(nums) {
		element := nums[index]
		if nums[index] != index+1 && nums[element-1] != element {
			nums[index], nums[element-1] = nums[element-1], nums[index]
		} else {
			index++
		}
	}
	ans := make([]int, 0)
	for index, element := range nums {
		if element != index+1 {
			ans = append(ans, index+1)
		}
	}
	return ans
}

