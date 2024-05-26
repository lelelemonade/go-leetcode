package problem279

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i < n+1; i++ {
		val := math.MaxInt

		for j := 1; j <= 100 && i-j*j >= 0; j++ {
			val = min(dp[i-j*j], val)
		}

		dp[i] = val + 1
	}

	return dp[n]
}

func Test279(t *testing.T) {
	assert.Equal(t, 2, numSquares(13))
	assert.Equal(t, 3, numSquares(12))
	assert.Equal(t, 4, numSquares(7168))
	assert.Equal(t, 1, numSquares(10000))
}
