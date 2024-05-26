package problem33

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test33(t *testing.T) {
	//testCase1 := []int{4, 5, 6, 7, 0, 1, 2}
	//assert.Equal(t, 4, search(testCase1, 0))
	//
	//testCase2 := []int{4, 5, 6, 7, 0, 1, 2}
	//assert.Equal(t, -1, search(testCase2, 3))
	//
	//testCase3 := []int{1}
	//assert.Equal(t, -1, search(testCase3, 0))
	//
	//testCase4 := []int{1, 3, 5}
	//assert.Equal(t, 0, search(testCase4, 1))
	//
	//testCase5 := []int{5, 1, 3}
	//assert.Equal(t, 0, search(testCase5, 5))

	testCase6 := []int{4, 5, 6, 7, 0, 1, 2}
	assert.Equal(t, 5, search(testCase6, 1))
}

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, start int, end int, target int) int {
	if start == end && nums[start] != target {
		return -1
	}

	if end-start == 1 {
		if nums[start] == target {
			return start
		} else if nums[end] == target {
			return end
		} else {
			return -1
		}
	}

	if nums[(end+start)/2] == target {
		return (end + start) / 2
	} else if nums[(end+start)/2] > target {
		if nums[start] == target {
			return start
		} else if nums[start] < nums[(end+start)/2] {
			return binarySearch(nums, (end+start)/2, end, target)
		} else {
			return binarySearch(nums, start, (end+start)/2, target)
		}
	} else {
		if nums[end] == target {
			return end
		} else if nums[start] < nums[(end+start)/2] {
			return binarySearch(nums, (end+start)/2, end, target)
		} else {
			return binarySearch(nums, start, (end+start)/2, target)
		}
	}
}
