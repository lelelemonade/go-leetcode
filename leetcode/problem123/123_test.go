package problem123

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test123(t *testing.T) {
	//assert.Equal(t, 6, maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	assert.Equal(t, 4, maxProfit([]int{1, 2, 3, 4, 5}))
}

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = -prices[0]
	dp[0][3] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i])
	}

	return dp[len(prices)-1][3]
}
