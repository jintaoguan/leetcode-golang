package main

import "fmt"

func main() {
	colors := []int{2, 0, 2, 1, 1, 0}
	sortColors(colors)
	fmt.Println(colors)
}

func sortColors(nums []int) {
	p0 := 0
	p1 := 0
	p2 := len(nums) - 1
	for p1 <= p2 {
		if nums[p1] == 1 {
			p1++
		} else if nums[p1] == 2 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p2--
		} else if nums[p1] == 0 {
			nums[p1], nums[p0] = nums[p0], nums[p1]
			p1++
			p0++
		}
	}
}

