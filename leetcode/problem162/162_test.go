package problem162

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test162(t *testing.T) {
	// assert.Equal(t,findPeakElement([]int{1,2}),1)
	assert.Equal(t, findPeakElement([]int{1, 2, 3, 1}), 2)
}

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if nums[i] > nums[i+1] {
				return i
			}
			continue
		}
		if i == len(nums)-1 {
			if nums[i] > nums[i-1] {
				return i
			}
			continue
		}
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}
