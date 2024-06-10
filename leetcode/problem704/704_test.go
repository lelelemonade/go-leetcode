package problem704

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test704(t *testing.T) {
	assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 2))
	assert.Equal(t, -1, search([]int{5}, -5))
}

func search(nums []int, target int) int {
	iter := len(nums) / 2
	left := 0
	right := len(nums) - 1

	for nums[iter] != target {
		if left == right {
			return -1
		}
		if nums[iter] > target {
			right = iter - 1
		} else {
			left = iter + 1
		}
		iter = (right + left) / 2
	}

	return iter
}

func searchRec(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if nums[len(nums)/2] == target {
		return len(nums) / 2
	} else if nums[len(nums)/2] <= target {
		return len(nums)/2 + 1 + searchRec(nums[len(nums)/2+1:], target)
	} else {
		return searchRec(nums[:len(nums)/2], target)
	}
}
