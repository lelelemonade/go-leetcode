package problem518

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test518(t *testing.T) {
	//assert.Equal(t, 4, change(5, []int{1, 2, 5}))
	assert.Equal(t, 0, change(3, []int{2}))
	assert.Equal(t, 1, change(10, []int{10}))
}

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coins[i]]
		}
	}

	return dp[len(dp)-1]
}
