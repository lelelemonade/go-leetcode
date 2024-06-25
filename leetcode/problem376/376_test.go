package problem376

import (
	"fmt"
	"testing"
)

func Test376(t *testing.T) {
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
}

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	topEndDp := make([]int, len(nums))
	topEndDp[0] = 1
	maxTopEnd := 1
	botEndDp := make([]int, len(nums))
	botEndDp[0] = 1
	maxBotEnd := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			topEndDp[i] = maxBotEnd + 1
			botEndDp[i] = topEndDp[i-1]
			maxTopEnd = max(maxTopEnd, topEndDp[i])
		} else if nums[i] == nums[i-1] {
			topEndDp[i] = botEndDp[i-1]
			botEndDp[i] = topEndDp[i-1]
		} else {
			botEndDp[i] = maxTopEnd + 1
			topEndDp[i] = botEndDp[i-1]
			maxBotEnd = max(maxBotEnd, botEndDp[i])
		}
	}

	return max(maxBotEnd, maxTopEnd)
}
