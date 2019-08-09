package main

import "fmt"

func main() {
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	last1 := 2
	last2 := 1
	ans := 0
	for i := 3; i <= n; i++ {
		ans = last1 + last2
		last2 = last1
		last1 = ans
	}

	return ans
}

