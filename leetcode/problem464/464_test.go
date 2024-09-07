package problem464

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal <= 0 {
		return true
	}
	if (maxChoosableInteger*(maxChoosableInteger+1))/2 < desiredTotal {
		return false
	}
	
	memo := make(map[int]bool)
	used := make([]bool, maxChoosableInteger+1)
	return canWinHelper(used, desiredTotal, memo)
}

func canWinHelper(used []bool, desiredTotal int, memo map[int]bool) bool {
	hash := format(used)
	if val, exists := memo[hash]; exists {
		return val
	}
	
	for i := 1; i < len(used); i++ {
		if !used[i] {
			used[i] = true
			if desiredTotal-i <= 0 || !canWinHelper(used, desiredTotal-i, memo) {
				memo[hash] = true
				used[i] = false
				return true
			}
			used[i] = false
		}
	}
	
	memo[hash] = false
	return false
}

func format(used []bool) int {
	hash := 0
	for i := 1; i < len(used); i++ {
		if used[i] {
			hash |= 1 << i
		}
	}
	return hash
}