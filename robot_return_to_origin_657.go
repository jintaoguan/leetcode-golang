package main

import (
	"fmt"
)

func main() {
	moves := "UD"
	fmt.Println(judgeCircle(moves))
}

func judgeCircle(moves string) bool {
	posX := 0
	posY := 0
	for _, ch := range moves {
		if ch == 'U' {
			posY++
		} else if ch == 'D' {
			posY--
		} else if ch == 'L' {
			posX--
		} else if ch == 'R' {
			posX++
		}
	}
	return posX == 0 && posY == 0
}

