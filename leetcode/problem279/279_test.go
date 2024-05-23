package problem279

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test279(t *testing.T) {
	assert.Equal(t, 3, numSquares(12))
	assert.Equal(t, 2, numSquares(13))
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i < len(dp); i++ {
		minSquare := math.MaxInt
		for j := 1; j*j <= i; j++ {
			minSquare = min(minSquare, dp[i-j*j]+1)
		}
		dp[i] = minSquare
	}

	return dp[n]
}
