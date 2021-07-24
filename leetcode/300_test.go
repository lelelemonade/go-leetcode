package leetcode

import (
	"fmt"
	"testing"
)

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	max := 0

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := i; j >= 0; j-- {
			if nums[j] < nums[i] && dp[j]+1 >= dp[i] {
				dp[i] = dp[j] + 1
			} else if nums[j] == nums[i] && dp[j] >= dp[i] {
				dp[i] = dp[j]
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

func TestLIS(t *testing.T) {
	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Println(lengthOfLIS(nums))
}
