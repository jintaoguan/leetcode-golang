package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(matrix))
}

func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		sum[i] = make([]int, m)
	}
	sum[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		sum[i][0] = sum[i-1][0] + grid[i][0]
	}
	for j := 1; j < m; j++ {
		sum[0][j] = sum[0][j-1] + grid[0][j]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			sum[i][j] = min(sum[i-1][j], sum[i][j-1]) + grid[i][j]
		}
	}
	return sum[n-1][m-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

