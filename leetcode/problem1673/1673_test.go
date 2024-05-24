package problem1673

func mostCompetitive(nums []int, k int) []int {
	stack := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && stack[len(stack)-1] > nums[i] && len(nums)-i+len(stack) > k {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}

	return stack[:k]
}
