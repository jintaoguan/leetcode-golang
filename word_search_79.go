package main

import "fmt"

func main() {
	board := [][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'},
	}
	fmt.Println(exist(board, "AAB"))
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			p := board[i][j]
			if p == word[0] {
				found := search(board, m, n, i, j, dirs, word, 0)
				if found {
					return true
				}
			}
		}
	}
	return false
}

func search(board [][]byte, m, n, x, y int, dirs [][]int, word string, index int) bool {
	var zerobyte byte
	if index == len(word)-1 {
		if board[x][y] == word[index] {
			return true
		}
		return false
	}
	if board[x][y] != word[index] {
		return false
	}
	original := board[x][y]
	board[x][y] = zerobyte
	for i := 0; i < 4; i++ {
		nx := x + dirs[i][0]
		ny := y + dirs[i][1]
		if nx < 0 || nx >= m || ny < 0 || ny >= n {
			continue
		}
		if board[nx][ny] == zerobyte {
			continue
		}
		found := search(board, m, n, nx, ny, dirs, word, index+1)
		if found {
			return true
		}
	}
	board[x][y] = original
	return false
}

