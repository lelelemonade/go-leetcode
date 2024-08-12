package problem14

func longestCommonPrefix(strs []string) string {
	result := ""
	finish := false

	for i := 0; ; i++ {
		var commonChar byte
		for j := 0; j < len(strs); j++ {
			if len(strs[j]) > i && commonChar == 0 {
				commonChar = strs[j][i]
			} else if len(strs[j]) > i && commonChar == strs[j][i] {

			} else {
				finish = true
			}
		}
		if !finish {
			result += string(commonChar)
		} else {
			break
		}
	}

	return result
}
