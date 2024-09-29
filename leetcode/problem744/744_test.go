package problem744

func nextGreatestLetter(letters []byte, target byte) byte {
	if target >= letters[len(letters)-1] {
		return letters[0]
	}
	left, right := 0, len(letters)-1
	for left <= right {
		mid := left + (right-left)/2
		if letters[mid] < target+1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return letters[left]
}
