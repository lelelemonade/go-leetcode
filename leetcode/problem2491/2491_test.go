package problem2491

import "slices"

func dividePlayers(skill []int) int64 {
	slices.Sort(skill)

	left, right := 0, len(skill)-1
	total := 0
	result := 0
	for left < right {
		if total == 0 {
			total = skill[left] + skill[right]
		} else if total != skill[left]+skill[right] {
			return -1
		}
		result += skill[left] * skill[right]
		left++
		right--
	}

	return int64(result)
}
