package problem1385

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	abs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	result := 0
	for _, v1 := range arr1 {
		shouldCount := true
		for _, v2 := range arr2 {
			if abs(v1, v2) <= d {
				shouldCount = false
				break
			}
		}
		if shouldCount {
			result++
		}
	}
	return result
}
