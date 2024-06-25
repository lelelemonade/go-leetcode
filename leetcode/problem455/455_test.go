package problem455

import "slices"

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	count := 0
	i := len(s) - 1

	for j := len(g) - 1; j >= 0; {
		if i < 0 || j < 0 {
			break
		}
		if s[i] >= g[j] {
			count++
			i--
		}
		j--

	}
	return count
}
