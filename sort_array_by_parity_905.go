package main

import (
	"fmt"
)

func main() {
	array := []int{3, 1, 2, 4}
	array2 := []int{0}
	fmt.Println(sortArrayByParity(array))
	fmt.Println(sortArrayByParity(array2))
}

func sortArrayByParity(A []int) []int {
	leftIndex := 0
	rightIndex := len(A) - 1
	for true {
		for leftIndex < len(A) && A[leftIndex]%2 == 0 {
			leftIndex++
		}
		for rightIndex >= 0 && A[rightIndex]%2 == 1 {
			rightIndex--
		}
		if leftIndex >= rightIndex {
			break
		}
		tmp := A[leftIndex]
		A[leftIndex] = A[rightIndex]
		A[rightIndex] = tmp
	}
	return A
}

