package problem459

import "strings"

func repeatedSubstringPattern(s string) bool {
	s2 := s + s
	return strings.Contains(s2[1:len(s2)-1], s)
}
