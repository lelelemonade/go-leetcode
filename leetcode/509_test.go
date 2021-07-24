package leetcode

import (
	"fmt"
	"testing"
)

func fib(n int) int {
	dp := make([]int, n+1)
	if n <= 1 {
		return n
	}
	dp[0] = 0
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func Test509(t *testing.T) {
	fmt.Println(fib(4))
}
