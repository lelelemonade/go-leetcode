package problem62

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test62(t *testing.T) {
	testCase62M := 3
	testCase62N := 7

	assert.Equal(t, 28, uniquePaths(testCase62M, testCase62N))
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				dp[i][j] = 1
			} else if j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
