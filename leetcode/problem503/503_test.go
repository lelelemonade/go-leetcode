package problem503

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test503(t *testing.T) {
	testCase1 := []int{1, 2, 3, 4, 3}
	testCase1Expect := []int{2, 3, 4, -1, 4}

	assert.Equal(t, testCase1Expect, nextGreaterElements(testCase1))
}

func nextGreaterElements(nums []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(nums))
	for i := 0; i < len(result); i++ {
		result[i] = -1
	}

	for i := 0; i < 2*len(nums)-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%len(nums)] {
			result[stack[len(stack)-1]] = nums[i%len(nums)]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i%len(nums))
	}

	return result
}
