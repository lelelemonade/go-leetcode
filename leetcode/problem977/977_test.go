package problem977

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	idx := len(nums) - 1
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			result[idx] = nums[left] * nums[left]
			left++
		} else {
			result[idx] = nums[right] * nums[right]
			right--
		}
		idx--
	}

	return result
}
