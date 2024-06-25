package problem47

import (
	"fmt"
	"slices"
	"testing"
)

func Test47(t *testing.T) {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 2, 3}))

}

func permuteUnique(nums []int) [][]int {
	results := make([][]int, 0)

	slices.Sort(nums)

	backtrack(make([]bool, len(nums)), []int{}, 0, nums, &results)

	return results
}

func backtrack(notChosen []bool, chosen []int, idx int, nums []int, results *[][]int) {
	if len(chosen) == len(nums) {
		tmp := make([]int, len(chosen))
		copy(tmp, chosen)
		*results = append(*results, tmp)
	}

	for i := 0; i < len(notChosen); i++ {
		if i != 0 && nums[i] == nums[i-1] && !notChosen[i-1] {
			continue
		}

		if !notChosen[i] {
			notChosen[i] = true
			backtrack(notChosen, append(chosen, nums[i]), idx+1, nums, results)
			notChosen[i] = false
		}
	}
}
