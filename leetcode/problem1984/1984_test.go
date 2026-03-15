package problem1984

import (
	"math"
	"slices"
)

func minimumDifference(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	slices.Sort(nums)

	minDiff := math.MaxInt

	for i := 0; i < len(nums) && i+k-1 < len(nums); i++ {
		minDiff = min(minDiff, nums[i+k-1]-nums[i])
	}

	return minDiff
}
