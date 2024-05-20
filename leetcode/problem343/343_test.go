package problem343

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test343(t *testing.T) {

	assert.Equal(t, 36, integerBreak(10))
	assert.Equal(t, 9, integerBreak(6))

}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i < n+1; i++ {
		maxInteger := 1
		for j := 0; j < i; j++ {
			maxInteger = max(maxInteger, max(dp[j], j)*(i-j))
		}

		dp[i] = maxInteger
	}

	return dp[len(dp)-1]
}
