package main

import (
	"fmt"
)

func main() {
	input := [][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}
	fmt.Println(flipAndInvertImage(input))
}

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		reverse(row)
		invert(row)
	}
	return A
}

func reverse(row []int) {
	for i := 0; i < len(row)/2; i++ {
		temp := row[i]
		row[i] = row[len(row)-i-1]
		row[len(row)-i-1] = temp
	}
}

func invert(row []int) {
	for i := 0; i < len(row); i++ {
		row[i] = row[i] ^ 1
	}
}

