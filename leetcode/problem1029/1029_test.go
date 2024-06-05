package problem1029

import "slices"

func twoCitySchedCost(costs [][]int) int {
	result := 0
	slices.SortFunc(costs, func(a []int, b []int) int {
		return abs(a[0]-a[1]) - abs(a[0]-a[1])
	})

	return result
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}
