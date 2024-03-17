package problem455

import "slices"

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		for j := len(g) - 1; j >= 0; {
			if s[i] >= g[len(g)-1] {
				count++
			}
			j--
		}
	}
	return count

}
