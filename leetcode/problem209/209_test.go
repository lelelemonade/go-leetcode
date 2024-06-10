package problem209

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test209(t *testing.T) {
	assert.Equal(t, 2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	assert.Equal(t, 1, minSubArrayLen(4, []int{1, 4, 4}))
	assert.Equal(t, 3, minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
}

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	left := 0
	result := math.MaxInt

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		if sum < target {
			continue
		}

		result = min(result, right-left+1)

		for sum > target {
			sum -= nums[left]
			left++

			if sum >= target {
				result = min(result, right-left+1)
			}
		}

	}

	if result == math.MaxInt {
		return 0
	}

	return result
}
