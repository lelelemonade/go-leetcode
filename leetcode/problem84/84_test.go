package problem84

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	stack := make([]int, 0)
	maxArea := 0

	for i, height := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > height {
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width -= stack[len(stack)-1] + 1
			}
			maxArea = max(maxArea, heights[left]*width)
		}
		stack = append(stack, i)
	}

	return maxArea
}

func Test84(t *testing.T){
	assert.Equal(t,10,largestRectangleArea([]int{2,1,5,6,2,3}))
}
