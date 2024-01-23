package problem2765

import (
	"fmt"
	"testing"
)

func alternatingSubarray(nums []int) int {
	maxLength := -1

	if len(nums) < 1 {
		return -1
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] == 1 {
			currentMaxLength := 2
			for j := i + 2; j < len(nums); j++ {
				if nums[j] == nums[(j-i)%2+i] {
					currentMaxLength += 1
				} else {
					break
				}
			}
			if maxLength < currentMaxLength {
				maxLength = currentMaxLength
			}
		}

	}

	return maxLength
}

func Test2765(t *testing.T) {
	fmt.Println(alternatingSubarray([]int{2, 3, 4, 3, 4}))
	fmt.Println(alternatingSubarray([]int{21, 3, 9}))
}
