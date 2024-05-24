package problem300

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test300(t *testing.T) {
	//assert.Equal(t, 4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	//assert.Equal(t, 4, lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	assert.Equal(t, 6, lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	for i := 0; i < len(dp); i++ {
		for j := i; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] == 0 {
			dp[i] = 1
		}
	}

	return slices.Max(dp)
}
