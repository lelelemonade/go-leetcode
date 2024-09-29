package problem907

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sumSubarrayMins(arr []int) int {
	stack := make([]int, 0)
	dp := make([]int,len(arr))
	result := 0

	for i, v := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > v {
			stack=stack[:len(stack)-1]
		}
		k := i+1
		if len(stack)>0{
			k=i-stack[len(stack)-1]
		}
		dp[i]=k*v
		if len(stack)>0{
			dp[i]+=dp[i-k]
		}
		result+=dp[i]
		result%=1000000007
		stack = append(stack, i)
	}
	return result
}

func Test907(t *testing.T){
	assert.Equal(t,17,sumSubarrayMins([]int{3,1,2,4}))
}
