package problem_backpack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBackpack01(t *testing.T) {
	assert.Equal(t, 12, backpack01([]int{5, 3, 2, 4, 1}, []int{3, 4, 2, 5, 1}, 12))
	assert.Equal(t, 35, backpack01([]int{15, 20, 30}, []int{1, 3, 4}, 4))
}

func backpack01(value []int, weight []int, capacity int) int {
	dp := make([][]int, len(value))
	for i := 0; i < len(value); i++ {
		dp[i] = make([]int, capacity+1)
	}

	for i := 0; i < len(value); i++ {
		for j := 0; j < capacity+1; j++ {
			if i == 0 {
				if weight[0] <= j {
					dp[0][j] = value[0]
				} else {
					dp[0][j] = 0
				}
				continue
			}

			if j-weight[i] >= 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(value)-1][capacity]
}
