package problem2683

func doesValidArrayExist(derived []int) bool {
	xorSum := 1
	for _, v := range derived {
		xorSum ^= v
	}
	return xorSum == 1
}
