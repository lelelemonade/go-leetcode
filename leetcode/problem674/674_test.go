package problem674

func findLengthOfLCIS(nums []int) int {
	result := 0
	current := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			current++
			result = max(result, current)
		} else {
			current = 1
		}
	}

	return result
}
