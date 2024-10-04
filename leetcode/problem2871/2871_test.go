package problem2871

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxSubarrays(nums []int) int {
	andSum := -1
	for _, v := range nums {
		andSum &= v
	}

	if andSum != 0 {
		return 1
	}
	result := 0
	i := 0
	for i < len(nums) {
		curAndSum := -1
		for {
			curAndSum &= nums[i]
			i++
			if curAndSum == 0 {
				result++
				break
			}
			if i == len(nums) {
				break
			}
		}
	}
	return result
}

func Test2871(t *testing.T) {
	assert.Equal(t, 3, maxSubarrays([]int{1, 0, 2, 0, 1, 2}))
}
