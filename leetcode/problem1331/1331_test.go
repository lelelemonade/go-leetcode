package problem1331

import "slices"

func arrayRankTransform(arr []int) []int {
	arrSorted := make([]int, len(arr))
	copy(arrSorted, arr)
	
	slices.Sort(arrSorted)

	ranks := make(map[int]int)
	for _, v := range arrSorted {
		if _, e := ranks[v]; !e {
			ranks[v] = len(ranks) + 1
		}
	}
	for i, v := range arr {
		arr[i] = ranks[v]
	}
	return arr
}
