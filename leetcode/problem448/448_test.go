package problem448

func findDisappearedNumbers(nums []int) []int {
	numSet := make(map[int]struct{})
	for _, v := range nums {
		numSet[v] = struct{}{}
	}
	result := []int{}
	for i := 1; i <= len(nums); i++ {
		if _, e := numSet[i]; !e {
			result = append(result, i)
		}
	}

	return result
}
