package problem450

import (
	"fmt"
	"sort"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestSlices(t *testing.T) {
	// Create a slice of strings
	strSlice := []string{"apple", "banana", "cherry", "date"}

	// Sort the slice of strings in descending order
	sort.SliceStable(strSlice, func(i, j int) bool {
		return strSlice[i] > strSlice[j]
	})

	// Print the sorted slice
	fmt.Println(strSlice)
}
