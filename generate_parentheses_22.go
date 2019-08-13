package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	result := make([]string, 0)
	current := make([]rune, n*2)
	generateParenthesisHelper(0, 0, current, &result)
	return result
}

func generateParenthesisHelper(index int, open int, current []rune, result *[]string) {
	if index >= len(current) {
		if open == 0 {
			*result = append(*result, string(current))
		}
		return
	}
	// add a open parenthesis
	if open < len(current)/2 {
		current[index] = '('
		generateParenthesisHelper(index+1, open+1, current, result)
	}
	if open > 0 {
		current[index] = ')'
		generateParenthesisHelper(index+1, open-1, current, result)
	}
}

