package problem875

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func minEatingSpeed(piles []int, h int) int {
	maxValue := slices.Max(piles)
	start := 1
	end := maxValue
	var result int
	for start <= end {
		i := (start + end) / 2
		time := 0
		for _, pile := range piles {
			time += (pile + i - 1) / i
		}
		if time <= h {
			result=i
			end = i - 1
		} else {
			start = i + 1
		}
	}

	return result
}

func Test875(t *testing.T) {
	assert.Equal(t,4,minEatingSpeed([]int{3,6,7,11},8))
	assert.Equal(t,30,minEatingSpeed([]int{30,11,23,4,20},5))
	assert.Equal(t,23,minEatingSpeed([]int{30,11,23,4,20},6))
	assert.Equal(t,2,minEatingSpeed([]int{312884470},312884469))
}