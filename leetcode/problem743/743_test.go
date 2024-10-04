package problem743

func networkDelayTime(times [][]int, n int, k int) int {
	alreadyKnown := make(map[int]int)
	alreadyKnown[k] = 0

	for {
		changed := false
		for _, time := range times {
			srcDelay, srcKnown := alreadyKnown[time[0]]
			dstDelay, dstKnown := alreadyKnown[time[1]]

			if srcKnown && (!dstKnown || srcDelay+time[2] < dstDelay) {
				alreadyKnown[time[1]] = srcDelay + time[2]
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	if len(alreadyKnown) != n {
		return -1
	}

	maxTime := 0
	for _, v := range alreadyKnown {
		maxTime = max(maxTime, v)
	}
	return maxTime
}
