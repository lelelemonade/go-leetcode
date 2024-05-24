package problem2363

import "slices"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	slices.SortFunc(items1, func(a []int, b []int) int {
		return a[0] - b[0]
	})
	slices.SortFunc(items2, func(a []int, b []int) int {
		return a[0] - b[0]
	})

	result := make([][]int, 0)
	for i, j := 0, 0; i < len(items1) || j < len(items2); {
		if i == len(items1) {
			for k := j; k < len(items2); k++ {
				result = append(result, []int{items2[k][0], items2[k][1]})
			}
			break
		} else if j == len(items2) {
			for k := i; k < len(items1); k++ {
				result = append(result, []int{items1[k][0], items1[k][1]})
			}
			break
		}

		if items1[i][0] == items2[j][0] {
			result = append(result, []int{items1[i][0], items1[i][1] + items2[j][1]})
			i++
			j++
		} else if items1[i][0] < items2[j][0] {
			result = append(result, []int{items1[i][0], items1[i][1]})
			i++
		} else {
			result = append(result, []int{items2[j][0], items2[j][1]})
			j++
		}
	}

	return result
}
