package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
}

func removeOuterParentheses(S string) string {
	sb := strings.Builder{}
	openParentheses := 0
	for _, ch := range S {
		if ch == '(' {
			if openParentheses > 0 {
				sb.WriteRune(ch)
			}
			openParentheses++
		} else if ch == ')' {
			openParentheses--
			if openParentheses > 0 {
				sb.WriteRune(ch)
			}
		}
	}
	return sb.String()
}

