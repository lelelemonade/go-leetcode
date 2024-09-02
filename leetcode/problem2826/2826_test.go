package problem2826

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minimumOperations(nums []int) int {
	f := [4]int{}
	for _, x := range nums {
		f[x]++
		f[2] = max(f[2], f[1])
		f[3] = max(f[3], f[2])
	}
	return len(nums) - max(f[1], f[2], f[3])
}

func Test2826(t *testing.T) {
	assert.Equal(t,2,minimumOperations([]int{2,3,1,2}))
}