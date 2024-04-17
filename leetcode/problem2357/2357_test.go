package problem2357

import (
	"fmt"
	"math"
	"testing"
)

func Test2357(t *testing.T) {
	fmt.Println(minimumOperations([]int{1, 5, 0, 3, 5}))
}

func minimumOperations(nums []int) int {
	minValue := math.MaxInt

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && nums[i] < minValue {
			minValue = nums[i]
		}
	}

	if minValue == math.MaxInt {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i] = nums[i] - minValue
		}
	}

	return minimumOperations(nums) + 1
}
