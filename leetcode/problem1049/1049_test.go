package problem1049

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1049(t *testing.T) {
	//assert.Equal(t, 1, lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
	//assert.Equal(t, 5, lastStoneWeightII([]int{31, 26, 33, 21, 40}))
	assert.Equal(t, 9, lastStoneWeightII([]int{21, 60, 61, 20, 31}))
}

func lastStoneWeightII(stones []int) int {

	sum := 0
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
	}

	dp := make([]int, sum/2+1)

	for i := 0; i < len(dp); i++ {
		if i < stones[0] {
			dp[i] = 0
		} else {
			dp[i] = stones[0]
		}
	}

	for i := 1; i < len(stones); i++ {
		for j := len(dp) - 1; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return sum - 2*dp[len(dp)-1]
}
