package problem42

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0

	trapWater := 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		if height[left] < height[right] {
			trapWater += leftMax - height[left]
			left++
		} else {
			trapWater += rightMax - height[right]
			right--
		}
	}
	return trapWater
}

func Test42(t *testing.T) {
	assert.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
