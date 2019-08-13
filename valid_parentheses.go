package main

import "fmt"

func main() {
	input := "()[]{}"
	fmt.Println(isValid(input))
}

func isValid(s string) bool {
	stack := make(map[int]rune, 0)
	stackTop := -1
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stackTop++
			stack[stackTop] = char
			continue
		}
		if stackTop < 0 {
			return false
		}
		if char == ')' && stack[stackTop] != '(' {
			return false
		} else if char == ']' && stack[stackTop] != '[' {
			return false
		} else if char == '}' && stack[stackTop] != '{' {
			return false
		}
		delete(stack, stackTop)
		stackTop--
	}
	return stackTop == -1
}

