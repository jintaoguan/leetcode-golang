package main

import (
	"fmt"
)

func main() {
	array := []int{5, 1, 5, 2, 5, 3, 5, 4}
	array2 := []int{1, 2, 3, 3}
	fmt.Println(repeatedNTimes1(array))
	fmt.Println(repeatedNTimes1(array2))
	fmt.Println(repeatedNTimes2(array))
	fmt.Println(repeatedNTimes2(array2))
}

func repeatedNTimes1(A []int) int {
	for k := 1; k <= 3; k++ {
		for i := 0; i < len(A)-k; i++ {
			if A[i] == A[i+k] {
				return A[i]
			}
		}
	}
	return -1
}

func repeatedNTimes2(A []int) int {
	counter := make(map[int]int)
	ret := A[0]
	for _, num := range A {
		if _, ok := counter[num]; !ok {
			counter[num] = 1
		} else {
			counter[num]++
		}
	}
	for k, v := range counter {
		if v >= len(A)/2 {
			ret = k
		}
	}
	return ret
}

