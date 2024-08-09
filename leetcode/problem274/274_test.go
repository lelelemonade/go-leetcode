package problem274

import (
	"slices"
)

func hIndex(citations []int) int {
	slices.Sort(citations)
	maxCitations := citations[len(citations)-1]
	currentIdx := len(citations) - 1

	for i := maxCitations; i >= 0; i-- {
		if len(citations)-currentIdx > i && citations[currentIdx] > i {
			return i
		}
		if citations[currentIdx]<i {
			currentIdx=i
		}
	}
}
