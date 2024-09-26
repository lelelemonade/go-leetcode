package problem2379

func minimumRecolors(blocks string, k int) int {
	blackCount := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'B' {
			blackCount++
		}
	}

	maxBlackCount := blackCount

	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'B' {
			blackCount++
		}
		if blocks[i-k] == 'B' {
			blackCount--
		}
		maxBlackCount = max(maxBlackCount, blackCount)
	}

	return k - maxBlackCount
}
