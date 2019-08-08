package main

import (
	"fmt"
)

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}

func numJewelsInStones(J string, S string) int {
	jewels := make(map[rune]bool)
	for _, letter := range J {
		if _, ok := jewels[letter]; !ok {
			jewels[letter] = true
		}
	}
	count := 0
	for _, letter := range S {
		if _, ok := jewels[letter]; ok {
			count++
		}
	}
	return count
}

