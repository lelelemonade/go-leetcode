package problem918

func maxSubarraySumCircular(nums []int) int {
	maxDp := make([]int, len(nums))
	maxDp[0] = nums[0]
	minDp := make([]int, len(nums))
	minDp[0] = nums[0]
	sum := nums[0]

	result := nums[0]
	minResult := nums[0]

	for i := 1; i < len(nums); i++ {
		maxDp[i] = max(maxDp[i-1]+nums[i], nums[i])
		minDp[i] = min(minDp[i-1]+nums[i], nums[i])
		result = max(result, maxDp[i])
		minResult = min(minResult, minDp[i])
		sum += nums[i]
	}

	if result < 0 {
		return result
	}

	return max(result, sum-minResult)
}
