package problem3097

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func minimumSubarrayLength(nums []int, k int) int {
	if k==0{
		return 1
	}
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, 0
	orSum := 0
	result := math.MaxInt
	for right < len(nums) {
		for right < len(nums) && orSum < k {
			orSum |= nums[right]
			right++
		}

		for left < len(nums) && orSum >= k {
			result = min(result, right-left)
			orSum &^= nums[left]
			left++
		}
	}
	if result == math.MaxInt{
		return -1
	}
	return result
}

func Test3097(t *testing.T) {
	// assert.Equal(t,1,minimumSubarrayLength([]int{1,2,3},2))
	// assert.Equal(t,3,minimumSubarrayLength([]int{2,1,8},10))
	// assert.Equal(t, 1, minimumSubarrayLength([]int{1, 2}, 0))
	assert.Equal(t, 3, minimumSubarrayLength([]int{1,2,32,21}, 55))
}
