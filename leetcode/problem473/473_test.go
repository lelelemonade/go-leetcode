package problem473

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test474(t *testing.T) {
	assert.Equal(t, 4, findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for j := 0; j < m+1; j++ {
		dp[j] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {
		zero, one := CountZeroAndOne(strs[i])
		for j := m; j >= zero; j-- {
			for k := n; k >= one; k-- {
				if i == 0 {
					if j >= zero && k >= one {
						dp[j][k] = 1
					} else {
						dp[j][k] = 0
					}
					continue
				}
				dp[j][k] = max(dp[j-zero][k-one]+1, dp[j][k])
			}
		}
	}

	return dp[m][n]
}

func CountZeroAndOne(str string) (zero, one int) {
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			zero++
		} else {
			one++
		}
	}
	return
}
