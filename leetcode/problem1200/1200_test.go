package problem1200

import (
	"math"
	"slices"
)

func minimumAbsDifference(arr []int) [][]int {
	ans := make([][]int, 0)

	slices.Sort(arr)

	minAbs := math.MaxInt64

	for i := 0; i+1 < len(arr); i++ {
		if arr[i+1]-arr[i] < minAbs {
			minAbs = arr[i+1] - arr[i]
			ans = make([][]int, 0)
			ans = append(ans, []int{arr[i], arr[i+1]})
		} else if arr[i+1]-arr[i] == minAbs {
			ans = append(ans, []int{arr[i], arr[i+1]})
		}
	}

	return ans
}
