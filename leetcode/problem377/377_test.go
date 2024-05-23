package problem377

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test377(t *testing.T) {
	//assert.Equal(t, 7, combinationSum4([]int{1, 2, 3}, 4))
	assert.Equal(t, 8, combinationSum4([]int{3, 1, 2, 4}, 4))
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
