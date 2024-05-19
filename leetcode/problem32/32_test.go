package problem32

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test32(t *testing.T) {
	testCase1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	testCase1Expect := 6

	assert.Equal(t, testCase1Expect, trap(testCase1))
}

func trap(height []int) int {
	rainCount := 0

	leftIter := 0
	rightIter := len(height) - 1
	leftMax := height[leftIter]
	rightMax := height[rightIter]

	for leftIter < rightIter {
		leftMax = max(height[leftIter], leftMax)
		rightMax = max(height[rightIter], rightMax)

		if leftMax < rightMax {
			if leftMax-height[leftIter] > 0 {
				rainCount += leftMax - height[leftIter]
			}
			leftIter++
		} else {
			if rightMax-height[rightIter] > 0 {
				rainCount += rightMax - height[rightIter]
			}
			rightIter--
		}
	}

	return rainCount
}
