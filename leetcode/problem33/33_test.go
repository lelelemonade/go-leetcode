package problem33

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test33(t *testing.T) {
	testCase0 := []int{3,1}
	assert.Equal(t, 1, search(testCase0, 1))

	testCase1 := []int{4, 5, 6, 7, 0, 1, 2}
	assert.Equal(t, 4, search(testCase1, 0))

	testCase2 := []int{4, 5, 6, 7, 0, 1, 2}
	assert.Equal(t, -1, search(testCase2, 3))

	testCase3 := []int{1}
	assert.Equal(t, -1, search(testCase3, 0))

	testCase4 := []int{1, 3, 5}
	assert.Equal(t, 0, search(testCase4, 1))

	testCase5 := []int{5, 1, 3}
	assert.Equal(t, 0, search(testCase5, 5))

	testCase6 := []int{4, 5, 6, 7, 0, 1, 2}
	assert.Equal(t, 5, search(testCase6, 1))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2
	for left <= right {
		midValue := nums[mid]

		if midValue == target {
			return mid
		} else if midValue >= nums[left] {
			if target < midValue && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target >midValue&&target<=nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		mid = (left + right) / 2
	}
	return -1
}
