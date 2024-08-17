package problem56

import "slices"

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	results := make([][]int, 0)
	results = append(results, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if results[len(results)-1][1] < intervals[i][0] {
			results = append(results, intervals[i])
		} else {
			results[len(results)-1] = []int{results[len(results)-1][0], max(results[len(results)-1][1],intervals[i][1])}
		}
	}

	return results
}
