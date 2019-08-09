package main

import "fmt"

func main() {
	array := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(array))
}

func majorityElement(nums []int) int {
	count := 1
	element := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != element {
			count--
			if count <= 0 {
				element = nums[i]
				count = 1
			}
		} else {
			count++
		}
	}
	return element
}

