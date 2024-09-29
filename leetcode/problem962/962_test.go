package problem962

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxWidthRamp(nums []int) int {
    stack := []int{}
    n := len(nums)
    for i := 0; i < n; i++ {
        if len(stack) == 0 || nums[i] < nums[stack[len(stack)-1]] {
            stack = append(stack, i)
        }
    }

    maxWidth := 0
    for j := n - 1; j >= 0; j-- {
        for len(stack) > 0 && nums[j] >= nums[stack[len(stack)-1]] {
            i := stack[len(stack)-1]
            width := j - i
            if width > maxWidth {
                maxWidth = width
            }
            stack = stack[:len(stack)-1]
        }
    }

    return maxWidth
}

func Test962(t *testing.T) {
	assert.Equal(t,7,maxWidthRamp([]int{9,8,1,0,1,9,4,0,4,1}))
}