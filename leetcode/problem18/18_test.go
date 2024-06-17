package problem18

import (
	"slices"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return make([][]int, 0)
	}

	slices.Sort(nums)

	result := make([][]int, 0)

	for a := 0; a < len(nums)-3; a++ {

		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		for b := a + 1; b < len(nums)-2; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			left := b + 1
			right := len(nums) - 1

			for left < right {
				if nums[a]+nums[b]+nums[left]+nums[right] < target {
					left++
				} else if nums[a]+nums[b]+nums[left]+nums[right] > target {
					right--
				} else {
					result = append(result, []int{nums[a], nums[b], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}
	}

	return result
}
