package problem242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sCount := make([]int, 26)

	for i := 0; i < len(s); i++ {
		sCount[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		sCount[t[i]-'a']--
		if sCount[t[i]-'a'] < 0 {
			return false
		}
	}

	for i := 0; i < len(t); i++ {
		if sCount[t[i]-'a'] != 0 {
			return false
		}
	}

	return true
}
