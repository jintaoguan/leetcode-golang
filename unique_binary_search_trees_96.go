package main

import "fmt"

func main() {
	fmt.Println(numTrees(3))
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			left := j
			right := i - j - 1
			sum += dp[left] * dp[right]
		}
		dp[i] = sum
	}
	return dp[n]
}

