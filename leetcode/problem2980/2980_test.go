package problem2980

func hasTrailingZeros(nums []int) bool {
	trailingZeroCount := 0
	for _, v := range nums {
		if v%2 == 0 {
			trailingZeroCount++
		}
	}
	return trailingZeroCount >= 2
}
