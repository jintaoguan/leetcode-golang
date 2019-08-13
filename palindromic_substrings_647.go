package main

import (
	"fmt"
)

func main() {
	s := "abc"
	fmt.Println(countSubstrings(s))
}

func countSubstrings(s string) int {
	count := 0
	count += len(s)
	for i := 0; i < len(s); i++ {
		for d := 0; d <= 1; d++ {
			low, high := i-d, i+1
			for low >= 0 && high < len(s) && s[low] == s[high] {
				count++
				low--
				high++
			}
		}
	}
	return count
}

