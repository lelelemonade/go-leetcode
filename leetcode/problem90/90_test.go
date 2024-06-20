package problem90

import (
	"fmt"
	"slices"
	"testing"
)

func Test90(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 2, 3}))
}

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)

	results := make([][]int, 0)

	backtrack(nums, []int{}, &results)

	return results
}

func backtrack(nums []int, chosen []int, results *[][]int) {
	result := make([]int, len(chosen))
	copy(result, chosen)
	*results = append(*results, result)

	for i := 0; i < len(nums); i++ {
		backtrack(nums[i+1:], append(chosen, nums[i]), results)
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
}
