package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-1; i++ {
		numsI := nums[i]
		j := i + 1
		k := len(nums) - 1
		if i > 0 && numsI == nums[i-1] {
			continue
		}
		for j != k {
			numsJ := nums[j]
			numsK := nums[k]
			if numsI+numsJ+numsK < 0 {
				j++
			} else if numsI+numsJ+numsK > 0 {
				k--
			} else {
				result = append(result, []int{numsI, numsJ, numsK})
				for j < k && numsJ == nums[j] {
					j++
				}
				for j < k && numsK == nums[k] {
					k--
				}
			}
		}
	}

	return result

}
