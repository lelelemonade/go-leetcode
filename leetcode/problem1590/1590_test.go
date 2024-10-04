package problem1590

func minSubarray(nums []int, p int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	mod := sum % p
	if mod == 0 {
		return 0
	}
	prefixSum := 0
	sumMap := map[int]int{0: -1}
	minLength := len(nums)

	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		curMod := (prefixSum + p) % p
		targetMod := (curMod - mod + p) % p
		if index, e := sumMap[targetMod]; e {
			minLength = min(minLength, i-index)
		}
		sumMap[curMod] = i
	}

	if minLength == len(nums) {
		return -1
	}
	return minLength
}
