package main

import "fmt"

func main() {
	input := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(input))
}

type position struct {
	x int
	y int
}

func numIslands(grid [][]byte) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	visited := make(map[position]bool)
	count := 0
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if _, ok := visited[position{i, j}]; ok {
				continue
			}
			if grid[i][j] == '1' {
				numIslandsHelper(grid, visited, position{i, j}, dirs)
				count++

			}

		}
	}
	return count
}

func numIslandsHelper(grid [][]byte, visited map[position]bool, pos position, dirs [][]int) {
	visited[pos] = true
	for i := 0; i < 4; i++ {
		nx := pos.x + dirs[i][0]
		ny := pos.y + dirs[i][1]
		if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) {
			continue
		}
		if _, ok := visited[position{nx, ny}]; !ok && grid[nx][ny] == '1' {
			numIslandsHelper(grid, visited, position{nx, ny}, dirs)
		}
	}
}

