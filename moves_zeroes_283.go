package main

import "fmt"

func main() {
	array := []int{0, 1, 0, 3, 12}
	moveZeroes(array)
	fmt.Println(array)
}

func moveZeroes(nums []int) {
	index := 0
	cur := 0
	for cur < len(nums) {
		if nums[cur] != 0 {
			if cur == index {
				cur++
				index++
				continue
			}
			tmp := nums[index]
			nums[index] = nums[cur]
			nums[cur] = tmp
			index++
		}
		cur++
	}
}

