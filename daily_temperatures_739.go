package main

import (
	"fmt"
)

func main() {
	temperatures := []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}
	fmt.Println(dailyTemperatures(temperatures))
}

func dailyTemperatures(T []int) []int {
	length := len(T)
	if length == 0 {
		return nil
	}
	waits := make([]int, length)
	stack := make(map[int][2]int)
	stackTop := -1
	for i := length - 1; i >= 0; i-- {
		temp := T[i]
		if stackTop < 0 {
			waits[i] = 0
		} else if T[i] < stack[stackTop][0] {
			waits[i] = 1
		} else {
			for T[i] >= stack[stackTop][0] && stackTop >= 0 {
				delete(stack, stackTop)
				stackTop--
			}
			if stackTop < 0 {
				waits[i] = 0
			} else {
				waits[i] = stack[stackTop][1] - i
			}
		}
		stackTop++
		stack[stackTop] = [2]int{temp, i}
	}
	return waits
}

