package problem96

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test96(t *testing.T) {
	assert.Equal(t, 5, numTrees(3))
	assert.Equal(t, 1, numTrees(1))
}

func numTrees(n int) int {
	if n <= 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] += 2 * dp[i-1]

		for j := 1; j <= i-1; j++ {
			dp[i] += dp[j-1] * dp[i-j-1]
		}
	}

	return dp[n-1]
}
