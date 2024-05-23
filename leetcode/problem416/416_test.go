package problem416

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test416(t *testing.T) {
	assert.Equal(t, true, canPartition([]int{1, 5, 11, 5}))
}

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)

	for j := 0; j < len(dp); j++ {
		if j < nums[0] {
			dp[j] = 0
		} else {
			dp[j] = nums[0]
		}
	}

	for i := 1; i < len(nums); i++ {
		for j := len(dp) - 1; j > 0; j-- {
			if j < nums[i] {
				dp[j] = dp[j]
			} else {
				dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
			}
		}
	}

	return dp[len(dp)-1] == target
}
