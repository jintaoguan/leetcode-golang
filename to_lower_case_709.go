package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toLowerCase1("Hello"))
	fmt.Println(toLowerCase2("Hello"))
}

func toLowerCase1(str string) string {
	sb := strings.Builder{}
	for _, ch := range str {
		if ch >= 'A' && ch <= 'Z' {
			sb.WriteRune(ch + 32)
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func toLowerCase2(str string) string {
	ret := make([]rune, len(str))
	for index, ch := range str {
		if ch >= 'A' && ch <= 'Z' {
			ret[index] = ch + 32
		} else {
			ret[index] = ch
		}
	}
	return string(ret)
}

