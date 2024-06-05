package problem1004

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1004(t *testing.T) {
	assert.Equal(t, 6, longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	assert.Equal(t, 10, longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
	assert.Equal(t, 4, longestOnes([]int{0, 0, 0, 1}, 4))
	assert.Equal(t, 3, longestOnes([]int{0, 0, 1, 1, 1, 0, 0}, 0))
}

func longestOnes(nums []int, k int) int {
	queue := make([]int, 0)
	firstOneIdx := 0

	longest := 0

	for i, v := range nums {
		if v == 0 {
			queue = append(queue, i)
			if len(queue) > k {
				firstOneIdx = queue[0] + 1
				queue = queue[1:]
				longest = max(longest, i-firstOneIdx+1)
			} else {
				longest = max(longest, i-firstOneIdx+1)
			}
		} else if v == 1 {
			longest = max(longest, i-firstOneIdx+1)
		}
	}

	return longest
}
