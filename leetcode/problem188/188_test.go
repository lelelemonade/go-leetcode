package problem188

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test188(t *testing.T) {
	//assert.Equal(t, 2, maxProfit(2, []int{2, 4, 1}))
	assert.Equal(t, 7, maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}

func maxProfit(k int, prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2*k)
	}

	for i := 0; i < len(dp[0]); i += 2 {
		dp[0][i] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if j == 0 {
				dp[i][j] = max(dp[i-1][j], -prices[i])
				continue
			}
			if j%2 == 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			}
		}
	}

	return dp[len(prices)-1][2*k-1]
}
