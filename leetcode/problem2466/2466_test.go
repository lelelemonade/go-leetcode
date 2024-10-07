package problem2466

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countGoodStrings(low int, high int, zero int, one int) int {
	mod := 1000000007
	dp := make([]int, high+1)
	dp[0] = 1

	for i := 1; i < high+1; i++ {
		if i >= zero {
			dp[i] = (dp[i] + dp[i-zero]) % mod
		}
		if i >= one {
			dp[i] = (dp[i] + dp[i-one]) % mod
		}
	}
	result := 0
	for i := low; i < high+1; i++ {
		result = (result + dp[i]) % mod
	}

	return result
}

func Test2446(t *testing.T) {
	assert.Equal(t, 8, countGoodStrings(3, 3, 1, 1))
}
