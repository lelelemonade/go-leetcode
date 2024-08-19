package problem650

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
    dp:=make([]int,n)
	dp[0]=0
	dp[1]=2

	for i:=2;i<n;i++{
		minDpN:=i+1

		for j:=2;j<=i/2;j++{
			if (i+1)%j==0{
				minDpN=min(minDpN,dp[j-1]+(i+1)/j)
			}
		}
		dp[i]=minDpN
	}

	return dp[n-1]
}

func Test650(t *testing.T) {
	// assert.Equal(t,3,minSteps(3))
	// assert.Equal(t,5,minSteps(6))
	assert.Equal(t,6,minSteps(9))
}