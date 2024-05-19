package problem2138

func divideString(s string, k int, fill byte) []string {
	result := make([]string, 0)

	for len(s) >= k {
		result = append(result, s[:k])
		s = s[k:]
	}

	if len(s) > 0 {
		for len(s) != k {
			s = s + string(fill)
		}
		result = append(result, s)
	}

	return result
}
