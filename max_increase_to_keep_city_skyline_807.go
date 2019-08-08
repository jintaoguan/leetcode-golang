package main

import "fmt"

func main() {
	grid := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	sum := 0
	m := len(grid)
	n := len(grid[0])
	gridnew := make([][]int, m)
	for i := 0; i < n; i++ {
		gridnew[i] = make([]int, n)
	}
	vView := generateVertialView(grid)
	hView := generateHorizontalView(grid)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			gridnew[i][j] = min(vView[j], hView[i])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum += gridnew[i][j] - grid[i][j]
		}
	}

	return sum
}

func generateVertialView(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	view := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				view[j] = grid[i][j]
			} else {
				view[j] = max(grid[i][j], view[j])
			}
		}
	}
	return view
}

func generateHorizontalView(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	view := make([]int, m)
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if j == 0 {
				view[i] = grid[i][j]
			} else {
				view[i] = max(grid[i][j], view[i])
			}
		}
	}
	return view
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

