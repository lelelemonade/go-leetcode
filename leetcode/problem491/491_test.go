package problem491

import (
	"fmt"
	"testing"
)

func Test491(t *testing.T) {
	fmt.Println(findSubsequences([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 1, 1, 1, 1}))
}

func findSubsequences(nums []int) [][]int {
	results := make([][]int, 0)

	backtrack(nums, []int{}, &results)

	return results
}

func backtrack(nums []int, chosen []int, results *[][]int) {
	if len(chosen) > 1 {
		result := make([]int, len(chosen))
		copy(result, chosen)
		*results = append(*results, result)
	}

	iteredBefore := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if iteredBefore[nums[i]] {
			continue
		}
		if len(chosen) == 0 || nums[i] >= chosen[len(chosen)-1] {
			backtrack(nums[i+1:], append(chosen, nums[i]), results)
			iteredBefore[nums[i]] = true
		}
	}
}
