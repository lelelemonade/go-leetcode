package problem3011

import (
	"github.com/stretchr/testify/assert"
	"math/bits"
	"testing"
)

func Test3011(t *testing.T) {
	assert.Equal(t, true, canSortArray([]int{8, 4, 2, 30, 15}))
}

func canSortArray(nums []int) bool {
	preMax := 0
	for i := 0; i < len(nums); {
		mx := nums[i]
		ones := bits.OnesCount(uint(mx))
		for ; i < len(nums) && bits.OnesCount(uint(nums[i])) == ones; i++ {
			if nums[i] < preMax {
				return false
			}
			mx = max(mx, nums[i])
		}
		preMax = mx
	}
	return true
}
