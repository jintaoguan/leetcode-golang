package main

import (
	"fmt"
)

func main() {
	array := []int{-4, -1, 0, 3, 10}
	// 16, 1, 0, 9, 100
	fmt.Println(sortedSquares(array))
}

func sortedSquares(A []int) []int {
	length := len(A)
	B := make([]int, length)

	for i, v := range A {
		A[i] = v * v
	}

	for i, j, k := 0, length-1, length-1; i <= j; k-- {
		if A[i] > A[j] {
			B[k] = A[i]
			i++
		} else {
			B[k] = A[j]
			j--
		}
	}

	return B
}

